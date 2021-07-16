package section1

/*
	插入排序算法
*/
func insertionSort(unorderedArray *[5]int32) {
	for i := range unorderedArray {
		preIndex := i - 1
		current := unorderedArray[i]
		for preIndex >= 0 && unorderedArray[preIndex] > current {
			unorderedArray[preIndex+1] = unorderedArray[preIndex]
			preIndex -= 1
		}
		unorderedArray[preIndex+1] = current
	}
}
