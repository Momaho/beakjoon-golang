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

	for ; m > 0; m-- {
		var i, j int
		var copy []int
		fmt.Fscan(reader, &i, &j)

		copy = append(copy, slice[i-1:j]...)
		for num := 0; num < len(copy); num++ {
			slice[i-1+num] = copy[len(copy)-1-num]
		}

	}
	for _, v := range slice {
		fmt.Print(v, " ")
	}
}
