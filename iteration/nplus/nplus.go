package main

import (
	"fmt"
)

func main() {
	var num, result int
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		result += i
	}
	fmt.Println(result)
}
