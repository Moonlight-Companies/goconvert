# Message and Convert Packages

This project consists of two main packages: `message` and `convert`. The `message` package provides utility functions for reading and converting values from `map[string]interface{}` structures, while the `convert` package offers a generic type conversion function.

## Features

- Generic `Read` function to retrieve and convert individual values from `map[string]interface{}`
- Generic `ReadList` function to retrieve and convert lists of values from `map[string]interface{}`
- Active conversions using the `convert.ConvertInto[T]` function
- Standalone `ConvertInto[T]` function for flexible type conversions

## Usage

### Message Package

#### Reading Individual Values

Use the `Read` function to retrieve and convert individual values from a map:

```go
package main

import (
    "fmt"
    "github.com/Moonlight-Companies/goconvert"
)

func main() {
    m := map[string]interface{}{
        "name": "John Doe",
        "age":  "30",
    }

    name, ok := message.Read[string](m, "name")
    if ok {
        fmt.Printf("Name: %s\n", name)
    }

    age, ok := message.Read[int](m, "age")
    if ok {
        fmt.Printf("Age: %d\n", age)
    }
}
```

#### Reading Lists

Use the `ReadList` function to retrieve and convert lists of values:

```go
package main

import (
    "fmt"
    "github.com/Moonlight-Companies/goconvert"
)

func main() {
    m := map[string]interface{}{
        "scores": []interface{}{85, 90, 95},
    }

    scores, ok := message.ReadList[int](m, "scores")
    if ok {
        fmt.Printf("Scores: %v\n", scores)
    }
}
```

### Convert Package

The `convert` package provides a `ConvertInto[T]` function that can be used independently of maps for flexible type conversions.

`ConvertInto[T]` will make a best effort to convert basic numeric types and String to or from other basic types or String

#### Using ConvertInto[T]

```go
package main

import (
    "fmt"
    "github.com/Moonlight-Companies/goconvert"
)

func main() {
    // Convert string to int
    strValue := "123"
    intValue, ok := convert.ConvertInto[int](strValue)
    if ok {
        fmt.Printf("Converted string to int: %d\n", intValue)
    }

    // Convert float to string
    floatValue := 3.14
    strValue, ok = convert.ConvertInto[string](floatValue)
    if ok {
        fmt.Printf("Converted float to string: %s\n", strValue)
    }

    // Convert bool to int
    boolValue := true
    intValue, ok = convert.ConvertInto[int](boolValue)
    if ok {
        fmt.Printf("Converted bool to int: %d\n", intValue)
    }
}
```

The `ConvertInto[T]` function supports conversions between various types, including:
- Numeric types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64)
- Strings
- Booleans
