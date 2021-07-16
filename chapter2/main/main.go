package main

import (
	"algorithm/chapter2/section1"
	"fmt"
)

func main() {
	var example = []int32{6, 7, 843, 43, 23}
	section1.InsertionSort(example)
	fmt.Println(example)
}
