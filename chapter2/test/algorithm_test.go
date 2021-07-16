package test

import (
	"algorithm/chapter2/section1"
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var array = []int32{3, 3, 5, 7, 10}
	section1.InsertionSort(array)
	fmt.Println(array)
}
