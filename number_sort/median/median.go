package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var readaer *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var input int
	var slice []int
	var sum int

	//입력
	for i := 0; i < 5; i++ {
		fmt.Fscan(readaer, &input)
		slice = append(slice, input)
	}

	//정렬
	sort.Ints(slice)

	//합산
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}

	fmt.Println(sum / 5)             //평균값
	fmt.Println(slice[len(slice)/2]) //정렬된 슬라이스 가운데 값

}
