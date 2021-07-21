package section1

/*
	选择排序
	@Param unorderedArray 未排序数组(切片)
*/
func SelectionSort(unorderedArray []int32) {
	for i := range unorderedArray {
		for j := i + 1; j < len(unorderedArray); j++ {
			if unorderedArray[i] > unorderedArray[j] {
				unorderedArray[i] = unorderedArray[i] ^ unorderedArray[j]
				unorderedArray[j] = unorderedArray[i] ^ unorderedArray[j]
				unorderedArray[i] = unorderedArray[i] ^ unorderedArray[j]
			}
		}
	}
}
