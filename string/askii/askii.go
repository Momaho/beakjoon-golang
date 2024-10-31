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
	askii := []byte(str)

	fmt.Print(askii[0])
}
