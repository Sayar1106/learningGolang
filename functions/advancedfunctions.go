package main

import "fmt"

func test(x int) {
	fmt.Println("hello!")
}

func test3(myFunc func(int) int) {
	 fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	return func() {fmt.Println(x)}
}

func main() {
	x := test
	x(50)

	test2 := func(x int) int {
		return x * -1
	}

	test3(test2)
	returnFunc("Hello")()
	y := returnFunc("goodbye")
	y()
}