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

	var num int

	fmt.Fscan(reader, &num)
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer)
	}
}
