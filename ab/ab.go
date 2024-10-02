package main

import (
	"fmt"
)

func main() {
	var a, b float32

	fmt.Scanln(&a, &b)
	fmt.Println(a / b)
}
