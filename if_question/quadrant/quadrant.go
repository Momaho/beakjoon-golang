package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	if x > 0 && y > 0 {
		fmt.Println(1)
	} else if x < 0 && y > 0 {
		fmt.Println(2)
	} else if x < 0 && y < 0 {
		fmt.Println(3)
	} else if x > 0 && y < 0 {
		fmt.Println(4)
	}
}
