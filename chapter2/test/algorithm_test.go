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
	array = []int32{3, 45, 5, 7, 10}
	section1.InsertionSortDesc(array)
	fmt.Println(array)
	fmt.Println("-----------二进制数加法运算---------------")
	array1 := []int{1, 0, 1, 1, 1}
	array2 := []int{1, 0, 1, 0, 1, 1}
	result := section1.BinaryCompute(array1, array2)
	fmt.Println(result)
	fmt.Println("-----------选择排序---------------")
	array = []int32{3, 45, 5, 7, 10}
	section1.SelectionSort(array)
	fmt.Println(array)
	fmt.Println("-----------归并排序---------------")
	array = []int32{3, 45, 5, 7, 10}
	array = section1.MergeSort(array)
	fmt.Println(array)
}
