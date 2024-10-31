package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func add(line int, num string) int {
	var sum int
	for i := 0; i < line; i++ {
		str, _ := strconv.Atoi(num[i : i+1])
		sum += str
	}
	return sum
}

func main() {
	var line int
	var num string
	fmt.Fscan(reader, &line)
	fmt.Fscan(reader, &num)

	fmt.Println(add(line, num))
}
