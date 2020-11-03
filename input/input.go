package main

import (
	"fmt";
	"os";
	"bufio";
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()
	// Convert string to proper type
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2020", 2020 - input)
}
