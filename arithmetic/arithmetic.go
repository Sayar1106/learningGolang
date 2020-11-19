package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int = 9
	var num2 int = 4
	answer := num1 % num2
	fmt.Printf("%d", answer)
	tan := math.Tan(float64(answer))
	fmt.Printf("%g", tan)

}
