package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var n, m int
	var slice []int

	fmt.Fscan(reader, &n, &m)

	for value := 0; value < n; value++ {
		slice = append(slice, value+1)
	}

	for ; m < 0; m-- {
		var i, j int

		fmt.Fscan(reader, &i, &j)
	}
}
