package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func lie_avg(m int, slice *[]int) float64 { //.평균
	var sum float64
	for _, v := range *slice {
		sum += float64(v) / float64(m) * 100
	}
	return sum / float64(len(*slice))
}

func max(slice *[]int) int { //.최대값
	var m int
	for _, v := range *slice {
		if m < v {
			m = v
		}
	}
	return m
}

func input(slice *[]int) { //.입력값
	var n, score int
	fmt.Fscan(reader, &n)

	for ; n > 0; n-- {
		fmt.Fscan(reader, &score)
		*slice = append(*slice, score)
	}
}

func main() {
	var slice []int
	var m int
	input(&slice)
	m = max(&slice)

	fmt.Println(lie_avg(m, &slice))
}
