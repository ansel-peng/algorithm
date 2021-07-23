package section1

/*
	归并排序
	unorderedArray 未排序数组
*/
func MergeSort(unorderedArray []int32) []int32 {
	length := len(unorderedArray)
	if length < 2 {
		return unorderedArray
	}
	middle := length / 2
	left := unorderedArray[0:middle]
	right := unorderedArray[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int32) []int32 {
	var result []int32
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}
