package section1

/*
	插入排序算法(升序)
    unorderedArray为无序数组(切片)
	现实举例:扑克牌
*/
func InsertionSort(unorderedArray []int32) {
	for i := range unorderedArray {
		preIndex := i - 1
		key := unorderedArray[i]
		for preIndex >= 0 && unorderedArray[preIndex] > key {
			unorderedArray[preIndex+1] = unorderedArray[preIndex]
			preIndex--
		}
		unorderedArray[preIndex+1] = key
	}
}

/*
	2.1-2 重写过程INSERTION-SORT,使之按非升序(而不是非降序)排序
	插入排序算法(降序)
	unorderedArray为无序数组(切片)
*/
func InsertionSortDesc(unorderedArray []int32) {
	maxIndex := len(unorderedArray) - 1
	for i := maxIndex; i >= 0; i-- {
		postIndex := i + 1
		key := unorderedArray[i]
		for postIndex <= maxIndex && unorderedArray[postIndex] > key {
			unorderedArray[postIndex-1] = unorderedArray[postIndex]
			postIndex++
		}
		unorderedArray[postIndex-1] = key
	}
}
