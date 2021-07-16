package test

import (
	"algorithm/chapter2/section1"
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	var array = []int32{3, 45, 5, 7, 10}
	fmt.Println("-----------升序插入排序---------------")
	section1.InsertionSort(array)
	fmt.Println(array)
	fmt.Println("-----------倒序插入排序---------------")
	array = []int32{323, 443, 433, 54, 65, 678}
	section1.InsertionSortDesc(array)
	fmt.Println(array)
}
