package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println((b % 100 % 10) * a)
	fmt.Println((b % 100 / 10) * a)
	fmt.Println((b / 100) * a)
	fmt.Println(a * b)
}
