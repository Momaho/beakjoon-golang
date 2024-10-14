package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	for {
		var a, b int
		_, err := fmt.Fscan(reader, &a, &b)
		if err == io.EOF {
			break
		}
		fmt.Fprintln(writer, a+b)
	}
}
