package main

import "fmt"
func add(x ,y int) (z1 , z2 int){
	defer fmt.Println("Hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("Before return")
	return
}

func main() {
	ans1, ans2 := add(5, 10)
	fmt.Println(ans1, ans2)
}
