package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, min, max int
	var slice []int

	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var input int
		fmt.Fscan(reader, &input)
		slice = append(slice, input)
	}
	for i, v := range slice {
		if i == 0 {
			min = v
			max = v
		} else if min > v {
			min = v
		} else if max < v {
			max = v
		}
	}
	fmt.Fprint(writer, min, max)
}
