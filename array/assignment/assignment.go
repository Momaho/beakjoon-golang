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
	var student = make([]int, 30)

	for i := 0; i < 28; i++ {
		var n int
		fmt.Fscan(reader, &n)
		student[n-1] += 1
	}

	for i, v := range student {
		if v != 1 {
			fmt.Fprintln(writer, i+1)
		}
	}
}
