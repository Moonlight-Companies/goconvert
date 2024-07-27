package message

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	m := map[string]interface{}{
		"string":  "hello",
		"int":     42,
		"float":   3.14,
		"bool":    true,
		"mixed":   "123",
		"invalid": struct{}{},
	}

	t.Run("ReadString", func(t *testing.T) {
		got, ok := Read[string](m, "string")
		if !ok || got != "hello" {
			t.Errorf("Read[string]() = %v, %v, want %v, true", got, ok, "hello")
		}
	})

	t.Run("ReadInt", func(t *testing.T) {
		got, ok := Read[int](m, "int")
		if !ok || got != 42 {
			t.Errorf("Read[int]() = %v, %v, want %v, true", got, ok, 42)
		}
	})

	t.Run("ReadFloat", func(t *testing.T) {
		got, ok := Read[float64](m, "float")
		if !ok || got != 3.14 {
			t.Errorf("Read[float64]() = %v, %v, want %v, true", got, ok, 3.14)
		}
	})

	t.Run("ReadBool", func(t *testing.T) {
		got, ok := Read[bool](m, "bool")
		if !ok || got != true {
			t.Errorf("Read[bool]() = %v, %v, want %v, true", got, ok, true)
		}
	})

	t.Run("ReadMixedAsString", func(t *testing.T) {
		got, ok := Read[string](m, "mixed")
		if !ok || got != "123" {
			t.Errorf("Read[string]() = %v, %v, want %v, true", got, ok, "123")
		}
	})

	t.Run("ReadMixedAsInt", func(t *testing.T) {
		got, ok := Read[int](m, "mixed")
		if !ok || got != 123 {
			t.Errorf("Read[int]() = %v, %v, want %v, true", got, ok, 123)
		}
	})

	t.Run("ReadNonexistent", func(t *testing.T) {
		_, ok := Read[string](m, "nonexistent")
		if ok {
			t.Errorf("Read[string]() ok = %v, want false", ok)
		}
	})

	t.Run("ReadInvalidType", func(t *testing.T) {
		_, ok := Read[string](m, "invalid")
		if ok {
			t.Errorf("Read[string]() ok = %v, want false", ok)
		}
	})
}

func TestReadList(t *testing.T) {
	m := map[string]interface{}{
		"ints":     []interface{}{1, 2, 3},
		"strings":  []interface{}{"a", "b", "c"},
		"mixed":    []interface{}{1, "2", 3.0},
		"invalid":  []interface{}{1, struct{}{}, 3},
		"notalist": "not a list",
	}

	t.Run("ReadIntList", func(t *testing.T) {
		got, ok := ReadList[int](m, "ints")
		want := []int{1, 2, 3}
		if !ok || !reflect.DeepEqual(got, want) {
			t.Errorf("ReadList[int]() = %v, %v, want %v, true", got, ok, want)
		}
	})

	t.Run("ReadStringList", func(t *testing.T) {
		got, ok := ReadList[string](m, "strings")
		want := []string{"a", "b", "c"}
		if !ok || !reflect.DeepEqual(got, want) {
			t.Errorf("ReadList[string]() = %v, %v, want %v, true", got, ok, want)
		}
	})

	t.Run("ReadMixedList", func(t *testing.T) {
		got, ok := ReadList[int](m, "mixed")
		want := []int{1, 2, 3}
		if !ok || !reflect.DeepEqual(got, want) {
			t.Errorf("ReadList[int]() = %v, %v, want %v, true", got, ok, want)
		}
	})

	t.Run("ReadInvalidList", func(t *testing.T) {
		got, ok := ReadList[int](m, "invalid")
		if ok {
			t.Errorf("ReadList[int]() = %v, %v, want [], false", got, ok)
		}
	})

	t.Run("ReadNonexistentList", func(t *testing.T) {
		_, ok := ReadList[int](m, "nonexistent")
		if ok {
			t.Errorf("ReadList[int]() ok = %v, want false", ok)
		}
	})

	t.Run("ReadNonList", func(t *testing.T) {
		_, ok := ReadList[int](m, "notalist")
		if ok {
			t.Errorf("ReadList[int]() ok = %v, want false", ok)
		}
	})
}
