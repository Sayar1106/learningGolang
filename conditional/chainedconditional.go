package main

import (
	"fmt"
)

func main() {
	x := 8
	val := (true || false) && !false || (x < 9)
	fmt.Println(val)
}
