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

	var n, x int
	var slice []int

	fmt.Fscan(reader, &n, &x)

	for i := 0; i < n; i++ {
		var input int

		fmt.Fscan(reader, &input)
		if input < x {
			slice = append(slice, input)
		}
	}

	for _, v := range slice {
		fmt.Fprintf(writer, "%d", v)
	}

}
