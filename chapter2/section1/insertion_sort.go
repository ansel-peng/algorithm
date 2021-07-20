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
	for i := range unorderedArray {
		preIndex := i - 1
		key := unorderedArray[i]
		for preIndex >= 0 && unorderedArray[preIndex] < key {
			unorderedArray[preIndex+1] = unorderedArray[preIndex]
			preIndex--
		}
		unorderedArray[preIndex+1] = key
	}
}

/*
	2.1-3 二进制计算
	array1 为第一个二进制数组， array2 为第二个二进制数组
*/
func BinaryCompute(a, b []int) []int {
	var result []int
	carryFlag := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) && i < len(b) {
			if (a[i] + b[i] + carryFlag) < 2 {
				result = append(result, a[i]+b[i]+carryFlag)
				carryFlag = 0
				continue
			} else {
				result = append(result, a[i]+b[i]+carryFlag-2)
				carryFlag = 1
				continue
			}
		}
		if i < len(a) {
			if (a[i] + carryFlag) < 2 {
				result = append(result, a[i]+carryFlag)
				carryFlag = 0
				continue
			} else {
				result = append(result, a[i]+carryFlag-2)
				carryFlag = 1
				continue
			}
		}
		if i < len(b) {
			if (b[i] + carryFlag) < 2 {
				result = append(result, b[i]+carryFlag)
				carryFlag = 0
				continue
			} else {
				result = append(result, b[i]+carryFlag-2)
				carryFlag = 1
				continue
			}
		}
	}
	if carryFlag == 1 {
		result = append(result, 1)
	}
	return result
}
