package section1

/*
	插入排序算法
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
