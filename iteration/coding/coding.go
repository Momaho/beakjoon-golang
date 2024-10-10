package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanln(&num)
	for i := num / 4; i > 0; i-- {
		fmt.Print("long ")
	}
	fmt.Print("int")
}
