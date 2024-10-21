package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func input(n, m int) []int {
	fmt.Fscan(reader, &n, &m)
	Baguni := make([]int, n)

	for lp := m; lp > 0; lp-- {
		var i, j, k int
		fmt.Fscan(reader, &i, &j, &k)
		for ; i <= j; i++ {
			Baguni[i-1] = k
		}
	}
	return Baguni
}

func main() {
	defer writer.Flush()

	var n, m int
	baguni := input(n, m)

	for _, v := range baguni {
		fmt.Fprint(writer, v, " ")
	}
}
