package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var s string

	s, _ = reader.ReadString('\n')

	split := strings.Fields(s)

	fmt.Print(len(split))

}
