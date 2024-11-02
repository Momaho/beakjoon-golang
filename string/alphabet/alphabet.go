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
	var s string

	fmt.Fscan(reader, &s)

	for alp := 97; alp >= 'a' && alp <= 'z'; alp++ {
		index := -1

		for i := 0; i < len(s); i++ {
			str := []byte(s[i : i+1])
			if str[0] == byte(alp) {
				index = i
				break
			}
		}
		fmt.Fprint(writer, index, " ")
	}

}
