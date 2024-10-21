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

	var max, max_n int
	var slice []int

	for i := 0; i < 9; i++ {
		var input int
		fmt.Fscan(reader, &input)
		slice = append(slice, input)
	}
	for i, v := range slice {
		if max < v {
			max = v
			max_n = i + 1
		}
	}
	fmt.Fprintf(writer, "%d\n%d", max, max_n)
}
