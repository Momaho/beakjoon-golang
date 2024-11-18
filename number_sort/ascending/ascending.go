package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, input int
	var slice []int
	fmt.Fscan(reader, &n) // n개의 수

	for i := 0; i < n; i++ { // 입력
		fmt.Fscan(reader, &input)
		slice = append(slice, input)
	}

	sort.Ints(slice) // 오름차순

	for i := 0; i < len(slice); i++ { // 출력
		fmt.Fprintln(writer, slice[i])
	}
	// sort.Sort(sort.Reverse(sort.IntSlice(slice))) // 내림차순

}
