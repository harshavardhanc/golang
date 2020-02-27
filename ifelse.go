package main

import "fmt"

func main() {
	var a, b int = 1, 2
	if a == b {
		fmt.Println("Values of a and b are equal", a, b)
	} else {
		fmt.Println("Values are not equal", a, b)
	}
}
