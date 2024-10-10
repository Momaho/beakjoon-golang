package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	if a == b && a == c {
		fmt.Println(10000 + a*1000)
	} else if a == b || a == c {
		fmt.Println(1000 + a*100)
	} else if b == c {
		fmt.Println(1000 + b*100)
	} else {
		if a > b {
			if a > c {
				fmt.Println(a * 100)
			} else {
				fmt.Println(c * 100)
			}
		} else {
			if b > c {
				fmt.Println(b * 100)
			} else {
				fmt.Println(c * 100)
			}
		}
	}
}
