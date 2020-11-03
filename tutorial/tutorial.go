package main

import "fmt"

func main() {
	var b1 bool
	var b2 int = 5
	var b3 float64 = 3.0
	var b4 = 56.23235
	b5 := "Sayar"
	fmt.Println(b1)
	fmt.Println(b2)
	// Printing type of var
	fmt.Printf("%T", b3)
	fmt.Printf("%T", b4)
	fmt.Println(b5)
	fmt.Printf("%T", b5)
}
