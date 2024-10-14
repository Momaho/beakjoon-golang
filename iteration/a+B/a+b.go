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

	var T int
	fmt.Fscan(reader, &T)
	for i := 1; i <= T; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, a+b)
	}
}
