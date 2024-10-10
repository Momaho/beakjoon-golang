package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}
