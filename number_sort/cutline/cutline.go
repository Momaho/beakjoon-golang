package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {

	var n, k, input int
	var slice []int

	fmt.Fscan(reader, &n, &k) // 응시자 수 , 상받는 사람 수 입력

	//학생들 점수 입력
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &input)
		slice = append(slice, input)
	}

	sort.Ints(slice) //정렬

	fmt.Println(slice[len(slice)-k]) //slice (len - 상받는 사람 의 수)
}
