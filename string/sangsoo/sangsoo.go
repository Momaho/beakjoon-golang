package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var redaer *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var s string
	var reverse string

	s, _ = redaer.ReadString('\n')

	for i := len(s); i > 0; i-- {
		reverse += s[i-1 : i]
	}
	split := strings.Fields(reverse)

	s1, _ := strconv.Atoi(split[0])
	s2, _ := strconv.Atoi(split[1])

	if s1 > s2 {
		fmt.Fprint(writer, s1)
	} else {
		fmt.Fprint(writer, s2)
	}
}
