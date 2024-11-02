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

	var s, r int
	var p string

	fmt.Fscan(reader, &s)

	for i := 0; i < s; i++ {
		fmt.Fscan(reader, &r, &p)
		for j := 0; j < len(p); j++ {
			for k := 0; k < r; k++ {
				fmt.Fprint(writer, p[j:j+1])
			}
		}
		fmt.Fprintln(writer)
	}
}
