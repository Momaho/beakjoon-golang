package main

import (
	"fmt"
)

func main() {
	var a, b, line int
	fmt.Scanln(&line)
	for i := 1; i <= line; i++ {
		fmt.Scanln(&a, &b)
		fmt.Println(a + b)
	}
}
