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

	for {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		if a == 0 && b == 0 {
			break
		} else {
			fmt.Fprintln(writer, a+b)
		}
	}
}
