package main

import "fmt"

func main() {
	a := make([]string, 4)
	fmt.Println("length of the slice", len(a))

	a = []string{"x", "y"}
	fmt.Println("Display:", a)

	a = append(a, "a", "b")
	fmt.Println("new:", a)

	fmt.Println("Print", a[:4])
}
