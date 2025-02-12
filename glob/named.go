package glob

import (
	"unicode"
	"unicode/utf8"
)

// MatchNamed reports whether text matches pattern and, if so,
// returns a map of named glob captures. A named glob begins
// with ':' and one or more letters (a-zA-Z). For example:
//
//	Pattern: "*/abc/:patterna/:patternb"
//	Text:    "/service/example_service/abc/1/2"
//
// returns true with map{"patterna": "1", "patternb": "2"}.
//
//	Pattern: "*/abc/user:userid"
//	Text:    "/service/example_service/abc/user123"
//
// returns true with map{"userid": "123"}.
//
// Matching is done case-insensitively.
func MatchNamed(pattern, text string) (bool, map[string]string, error) {
	caps := make(map[string]string)
	ok, finalCaps, err := matchHelper(pattern, text, caps)
	if err != nil {
		return false, nil, err
	}
	return ok, finalCaps, nil
}

// matchHelper is a recursive matcher that processes pattern p and text t.
// It carries along a map of captures (for named globs). It returns a boolean
// indicating success, the (possibly updated) capture map, and an error.
func matchHelper(p, t string, caps map[string]string) (bool, map[string]string, error) {
	// When the pattern is finished, the text must also be exhausted.
	if p == "" {
		return t == "", caps, nil
	}
	// Decode the next rune from the pattern.
	r, size := utf8.DecodeRuneInString(p)
	switch r {
	case '*':
		// Skip over any consecutive '*' characters.
		for {
			r2, s2 := utf8.DecodeRuneInString(p)
			if r2 != '*' {
				break
			}
			p = p[s2:]
			// A trailing '*' matches all the rest.
			if p == "" {
				return true, caps, nil
			}
		}
		// Try every possible split of t.
		for i := 0; i <= len(t); {
			ok, newCaps, err := matchHelper(p, t[i:], copyMap(caps))
			if err != nil {
				return false, nil, err
			}
			if ok {
				return true, newCaps, nil
			}
			if i == len(t) {
				break
			}
			_, sz := utf8.DecodeRuneInString(t[i:])
			i += sz
		}
		return false, nil, nil

	case '?':
		// '?' must match a single rune.
		if t == "" {
			return false, nil, nil
		}
		_, tsz := utf8.DecodeRuneInString(t)
		return matchHelper(p[size:], t[tsz:], caps)

	case '[':
		// Delegate character class matching.
		if t == "" {
			return false, nil, nil
		}
		tr, tsz := utf8.DecodeRuneInString(t)
		ok, rest, err := matchCharClass(p, tr)
		if err != nil || !ok {
			return false, nil, err
		}
		return matchHelper(rest, t[tsz:], caps)

	case '\\':
		// '\' escapes the next rune.
		p = p[size:]
		if p == "" {
			return false, nil, ErrBadPattern
		}
		rEsc, escSize := utf8.DecodeRuneInString(p)
		if t == "" {
			return false, nil, nil
		}
		tr, tsz := utf8.DecodeRuneInString(t)
		if unicode.ToUpper(rEsc) != unicode.ToUpper(tr) {
			return false, nil, nil
		}
		return matchHelper(p[escSize:], t[tsz:], caps)

	case ':':
		// Check if a named glob follows.
		temp := p[size:]
		nameRunes := []rune{}
		totalConsumed := 0
		for len(temp) > 0 {
			rTemp, sz := utf8.DecodeRuneInString(temp)
			if unicode.IsLetter(rTemp) {
				nameRunes = append(nameRunes, rTemp)
				totalConsumed += sz
				temp = temp[sz:]
			} else {
				break
			}
		}
		// If at least one letter was consumed, treat it as a named glob.
		if len(nameRunes) > 0 {
			name := string(nameRunes)
			pRest := p[size+totalConsumed:]
			// Try every possible split of t for the capture.
			for i := 0; i <= len(t); {
				newCaps := copyMap(caps)
				newCaps[name] = t[:i]
				ok, finalCaps, err := matchHelper(pRest, t[i:], newCaps)
				if err != nil {
					return false, nil, err
				}
				if ok {
					return true, finalCaps, nil
				}
				if i == len(t) {
					break
				}
				_, sz := utf8.DecodeRuneInString(t[i:])
				i += sz
			}
			return false, nil, nil
		}
		// If not a valid named glob, treat ':' as a literal.
		if t == "" {
			return false, nil, nil
		}
		tr, tsz := utf8.DecodeRuneInString(t)
		if unicode.ToUpper(r) != unicode.ToUpper(tr) {
			return false, nil, nil
		}
		return matchHelper(p[size:], t[tsz:], caps)

	default:
		// Match a literal character.
		if t == "" {
			return false, nil, nil
		}
		tr, tsz := utf8.DecodeRuneInString(t)
		if unicode.ToUpper(r) != unicode.ToUpper(tr) {
			return false, nil, nil
		}
		return matchHelper(p[size:], t[tsz:], caps)
	}
}

// matchCharClass processes a character class from p (which is expected
// to begin with '[') and tests whether ch is in that class. It returns a
// boolean indicating the match, the remaining pattern after the class, and
// an error if the syntax is bad.
func matchCharClass(p string, ch rune) (bool, string, error) {
	if p == "" || p[0] != '[' {
		return false, "", ErrBadPattern
	}
	// Skip the opening '['.
	p = p[1:]
	negated := false
	if p != "" {
		r, sz := utf8.DecodeRuneInString(p)
		if r == '^' {
			negated = true
			p = p[sz:]
		}
	}
	matched := false
	nrange := 0
	// Process character ranges until the closing ']'
	for {
		if p == "" {
			return false, "", ErrBadPattern
		}
		r, sz := utf8.DecodeRuneInString(p)
		// A closing ']' after at least one range ends the class.
		if r == ']' && nrange > 0 {
			p = p[sz:]
			break
		}
		var lo rune
		if r == '\\' {
			p = p[sz:]
			if p == "" {
				return false, "", ErrBadPattern
			}
			lo, sz = utf8.DecodeRuneInString(p)
			p = p[sz:]
		} else {
			lo = r
			p = p[sz:]
		}
		hi := lo
		// If a '-' follows, then treat as a range.
		if p != "" {
			r2, sz2 := utf8.DecodeRuneInString(p)
			if r2 == '-' {
				p = p[sz2:]
				if p == "" {
					return false, "", ErrBadPattern
				}
				r3, sz3 := utf8.DecodeRuneInString(p)
				if r3 == '\\' {
					p = p[sz3:]
					if p == "" {
						return false, "", ErrBadPattern
					}
					r3, sz3 = utf8.DecodeRuneInString(p)
					p = p[sz3:]
				} else {
					p = p[sz3:]
				}
				hi = r3
				if hi < lo {
					return false, "", ErrBadPattern
				}
			}
		}
		nrange++
		// If ch falls within [lo,hi] (case-insensitively), mark it.
		if unicode.ToUpper(lo) <= unicode.ToUpper(ch) && unicode.ToUpper(ch) <= unicode.ToUpper(hi) {
			matched = true
		}
	}
	if negated {
		matched = !matched
	}
	return matched, p, nil
}

// copyMap returns a shallow copy of the map.
func copyMap(m map[string]string) map[string]string {
	nm := make(map[string]string, len(m))
	for k, v := range m {
		nm[k] = v
	}
	return nm
}
