package main

import (
	"fmt"
)

func main() {
	var h, m int
	fmt.Scanln(&h, &m)

	m -= 45
	if m >= 0 {
		fmt.Println(h, m)
	} else if m < 0 {
		h -= 1
		if h < 0 {
			fmt.Println(23, m+60)
		} else {
			fmt.Println(h, m+60)
		}
	}

}
