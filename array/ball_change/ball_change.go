package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func change(baguni []int, m int) {
	for i := 0; i < len(baguni); i++ {
		baguni[i] = i + 1
	}

	for lp := m; lp > 0; lp-- {
		var i, j, temp int
		fmt.Fscan(reader, &i, &j)
		temp = baguni[i-1]
		baguni[i-1] = baguni[j-1]
		baguni[j-1] = temp
	}
}

func main() {
	defer writer.Flush()
	var n, m int

	fmt.Fscan(reader, &n, &m)
	baguni := make([]int, n)

	change(baguni, m)

	for _, v := range baguni {
		fmt.Fprint(writer, v, " ")
	}
}
