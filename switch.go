package main

import "fmt"

func main() {
	marks := 60
	switch marks {
	case 60:
		fmt.Println("B")
	case 70:
		fmt.Println("A")
	case 80:
		fmt.Println("A+")
	case 90, 95:
		fmt.Println("A++")
	}

}
