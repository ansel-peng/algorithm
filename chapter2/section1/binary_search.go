package section1

/*
	BinarySearch 二分查找
	array 已排序号数组
	target 寻找目标
*/
func BinarySearch(array []int, target int) int {
	if array == nil {
		return -1
	}
	start := 0
	end := len(array) - 1
	for start <= end {
		mid := start + (end-start)/2
		if array[mid] == target {
			return mid
		} else if target < array[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
