package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var str string
	var n int

	fmt.Fscan(reader, &str)
	fmt.Fscan(reader, &n)

	fmt.Printf("%c", str[n-1])
}
