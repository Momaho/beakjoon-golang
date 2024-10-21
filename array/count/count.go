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
	var a, value, num int
	var slice []int
	fmt.Fscan(reader, &a)

	for i := 0; i < a; i++ {
		fmt.Fscan(reader, &value)
		slice = append(slice, value)
	}
	a = 0
	fmt.Fscan(reader, &num)

	for _, v := range slice {
		if num == v {
			a++
		}
	}
	fmt.Fprintln(writer, a)
}
