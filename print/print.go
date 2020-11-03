package main

import "fmt"
func main() {
	// var type
	fmt.Printf("%T\n", 56.1)
	//Binary
	fmt.Printf("%b\n", 1203)
	// Octal
	fmt.Printf("%o\n", 1203)
	// Hexa
	fmt.Printf("%X\n", 1204)
	// Scientific notation
	fmt.Printf("%e\n", 1.439039095)
	// Large exponents
	fmt.Printf("%g\n", 1.4490923)
	// String
	fmt.Printf("%s\n", "Sayar")
	// Double quotes string
	fmt.Printf("%q\n", "Sayar")
	// Padding
	fmt.Printf("%07d\n", 103)
	// Store print
	var string = fmt.Sprintf("Sayar")
	fmt.Println(string)
}
