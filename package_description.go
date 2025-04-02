/*
Package stable provides a powerful and flexible library for creating ASCII tables in Go.
It supports a wide variety of input data types and offers extensive customization options
for table formatting and display.

Data Source Support:
  - Structs and struct arrays
  - JSON encoded byte arrays
  - String-interface maps and map arrays
  - CSV encoded strings
  - Custom row-by-row values

Key Features:
  - Comprehensive type support for various data structures
  - Customizable value and header alignment
  - Flexible formatting options with printf-style syntax
  - Character limiting with configurable truncation
  - Customizable border styles and table appearance
  - Field hiding and visibility control
  - Summary calculations (sum, average, min, max)
  - Automatic column width calculation
  - Support for custom captions and headers

Example Usage:

	table := stable.New("Users")
	table.AddField("Name")
	table.AddField("Age")
	table.AddField("Email")
	table.Row("John Doe", 30, "john@example.com")
	fmt.Println(table)

For more examples and detailed documentation, visit:

	https://github.com/ecoshub/stable
*/
package stable
