package section1

/*
	BubbleSort
	冒泡排序
	unorderedArray 未排序数组
*/
func BubbleSort(unorderedArray []int32) {
	for i := range unorderedArray {
		for j := 0; j < len(unorderedArray)-i-1; j++ {
			if unorderedArray[j] > unorderedArray[j+1] {
				unorderedArray[j], unorderedArray[j+1] = unorderedArray[j+1], unorderedArray[j]
			}
		}
	}
}
