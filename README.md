# Stable - Elegant ASCII Tables for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/ecoshub/stable)](https://goreportcard.com/report/github.com/ecoshub/stable)
[![GoDoc](https://godoc.org/github.com/ecoshub/stable?status.svg)](https://godoc.org/github.com/ecoshub/stable)
[![License](https://img.shields.io/github/license/ecoshub/stable)](https://github.com/ecoshub/stable/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ecoshub/stable)](https://github.com/ecoshub/stable/blob/master/go.mod)
[![Release](https://img.shields.io/github/v/release/ecoshub/stable)](https://github.com/ecoshub/stable/releases)

Stable is a powerful and flexible Go library for creating beautiful ASCII tables. It supports a wide variety of input data types and offers extensive customization options for table formatting and display.

## ✨ Features

- **Multiple Data Source Support**
  - Structs and struct arrays
  - JSON encoded byte arrays
  - String-interface maps and map arrays
  - CSV encoded strings
  - Custom row-by-row values

- **Rich Formatting Options**
  - Customizable value and header alignment
  - Flexible printf-style formatting
  - Character limiting with configurable truncation
  - Customizable border styles
  - Field hiding and visibility control
  - Summary calculations (sum, average, min, max)
  - Automatic column width calculation
  - Custom captions and headers

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/ecoshub/stable"
)

func main() {
    // Create a new table with caption
    table := stable.New("Users")

    // Add fields
    table.AddField("Name")
    table.AddField("Age")
    table.AddField("Email")

    // Add rows
    table.Row("John Doe", 30, "john@example.com")
    table.Row("Jane Smith", 25, "jane@example.com")

    // Print the table
    fmt.Println(table)

    // Output:
	// +-------------------------------------------+
	// |                   Users                   |
	// |-------------------------------------------|
	// |     Name     |  Age  |        Email       |
	// |--------------+-------+--------------------|
	// |  John Doe    |  30   |  john@example.com  |
	// |  Jane Smith  |  25   |  jane@example.com  |
	// +--------------+-------+--------------------+
}
```

## 📚 Examples

### From Struct Array
```go
type Person struct {
    Name      string
    Age       int
    Height    float64
    Available bool
}

persons := []Person{
    {Name: "Ruby Cohen", Age: 30, Height: 1.80, Available: true},
    {Name: "Bethany Parsons", Age: 29, Height: 1.58, Available: false},
    {Name: "Ronnie Rodriguez", Age: 28, Height: 1.78, Available: true},
    {Name: "Rosa Daniels", Age: 31, Height: 1.80, Available: true},
}

// Convert to table
t, err := stable.ToTable(persons)
if err != nil {
    fmt.Println(err)
    return
}

t.SetCaption("Customers")
fmt.Println(t)
// Output:
// +-----------------------------------------------------+
// |                      Customers                      |
// |-----------------------------------------------------|
// |        Name        |  Age  |  Height  |  Available  |
// |--------------------+-------+----------+-------------|
// |  Ruby Cohen        |  30   |  1.8     |  true       |
// |  Bethany Parsons   |  29   |  1.58    |  false      |
// |  Ronnie Rodriguez  |  28   |  1.78    |  true       |
// |  Rosa Daniels      |  31   |  1.8     |  true       |
// +--------------------+-------+----------+-------------+
```

### Custom Formatting
```go
table := stable.New("Benchmark Results")

// Add fields with custom options
table.AddField("File")
table.AddFieldWithOptions("Size (KB)", &stable.Options{
    Format:    "%.2f",
    Alignment: stable.AlignmentRight,
})
table.AddFieldWithOptions("Time (ms)", &stable.Options{
    Format:    "%.3f",
    Alignment: stable.AlignmentCenter,
})

// Add rows
table.Row("/var/log/sys.log", 12.515, 32.223)
table.Row("/var/log/app.log", 85.521, 311.742)

fmt.Println(table)
// Output:
// +--------------------------------------------------------+
// |                    Benchmark Results                   |
// |--------------------------------------------------------|
// |        File        |    Size (KB)    |    Time (ms)    |
// |--------------------+-----------------+-----------------|
// |  /var/log/sys.log  |          12.52  |      32.223     |
// |  /var/log/app.log  |          85.52  |     311.742     |
// +--------------------+-----------------+-----------------+
```

## 📖 Documentation

For detailed documentation and more examples, visit:
- [GoDoc](https://godoc.org/github.com/ecoshub/stable)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 🙏 Acknowledgments

- Inspired by various table formatting libraries
- Built with ❤️ for the Go community
