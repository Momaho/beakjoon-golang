package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var n int
	li := []int{}

	//입력
	fmt.Fscan(reader, &n)
	for ; n > 0; n-- {
		var e int
		fmt.Fscan(reader, &e)
		li = append(li, e)
	}

	//정렬
	sort.Ints(li)
	//출력
	for i := 0; i < len(li); i++ {
		fmt.Println(li[i])
	}

}
