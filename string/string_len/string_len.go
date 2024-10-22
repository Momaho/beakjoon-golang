package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var str string

	fmt.Fscan(reader, &str)
	fmt.Print(len(str))
}
