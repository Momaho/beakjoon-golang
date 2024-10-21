package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	var array [10]int
	var n, count int

	for i, _ := range array {
		fmt.Fscan(reader, &n)
		array[i] = n % 42
		count += 1
		for j, _ := range array {
			if j == i {
				break
			} else if array[i] == array[j] {
				count -= 1
				break
			}
		}
	}
	fmt.Println(count)

}
