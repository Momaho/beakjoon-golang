package main

import (
	"fmt"
)

func main() {
	var h, m, c int

	fmt.Scanln(&h, &m)
	fmt.Scanln(&c)

	if m+c < 60 {
		fmt.Println(h, m+c)
	} else {
		if (h + (m+c)/60) >= 24 {
			fmt.Println(h+((m+c)/60)-24, (m+c)%60)
		} else {
			fmt.Println(h+((m+c)/60), (m+c)%60)
		}
	}
}
