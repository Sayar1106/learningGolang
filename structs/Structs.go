package main

import "fmt"

type Point struct {
	x int32
	y int32
	isOnGrid bool	
}

func main(){
	var p1 Point = Point{1, 2, false}
	var p2 Point = Point{-5, 7, true}
	p1.x = 7
	fmt.Println(p1.y)
	fmt.Println(p2.x)
}