package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n string
	var slice []int

	fmt.Fscan(reader, &n)

	for i := 0; i < len(n); i++ { // 입력
		num, _ := strconv.Atoi(n[i : i+1])
		slice = append(slice, num)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(slice))) // 내림차순

	for i := 0; i < len(slice); i++ { //출력
		fmt.Fprint(writer, slice[i])
	}

}
