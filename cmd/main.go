package main

type Person struct {
	Age    int     `table:"Age"`
	Height float64 `table:"Height"`
	Name   string  `table:"Name"`
	Male   bool    `table:"Male"`
}

func main() {
	// anonymus()
	test()
	// printStructArray()
	// printStruct()
}
