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
	fmt.Println("-----------冒泡排序---------------")
	array = []int32{3, 45, 5, 7, 10}
	section1.BubbleSort(array)
	fmt.Println(array)
	fmt.Println("-----------二分查找---------------")
	sortArray := []int{3, 5, 7, 10, 45}
	index := section1.BinarySearch(sortArray, 7)
	fmt.Println(index)
	fmt.Println("-----------逆序对---------------")
	testArray := []int{45, 5, 7, 10, 3}
	count := section1.ReversePairs(testArray)
	fmt.Println(count)

}
