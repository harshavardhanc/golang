package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Println("second element is:", a[1])

	var b [5]int
	b = [5]int{10, 20, 30, 40}
	fmt.Println("All elements:", b)
}
