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

	for i := num; i > 0; i-- {
		for j := num; j > 0; j-- {
			if j > num+1-i {
				fmt.Fprint(writer, " ")
			} else {
				fmt.Fprint(writer, "*")
			}
		}
		fmt.Fprintln(writer)
	}
}
