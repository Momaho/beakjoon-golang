package main

import (
	"fmt"
)

func main() {
	var x, n, a, b, result int
	fmt.Scanln(&x)
	fmt.Scanln(&n)
	for i := n; i > 0; i-- {
		fmt.Scanln(&a, &b)
		result += a * b
	}
	if x == result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
