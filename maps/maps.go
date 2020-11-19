package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	fmt.Println(mp["apple"])
	mp["apple"] = 900
	mp["tim"] = 900
	delete(mp, "apple")
	val, ok := mp["sayar"]
	fmt.Println(val, ok)
	fmt.Println(mp)
}
