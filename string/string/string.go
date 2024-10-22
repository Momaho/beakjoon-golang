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

	var t int
	var str string

	fmt.Fscan(reader, &t)
	for ; t > 0; t-- {
		fmt.Fscan(reader, &str)
		fmt.Fprintf(writer, "%c%c ", str[0], str[len(str)-1])
	}
}
