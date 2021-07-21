package section1

/*
	2.1-4 考虑把两个n位二进制整数加起来的问题，这两个整数分别存储在两个n元数组A和B中。这两个整数的和应按二进制形式存储在一个
          (n+1)元数组C中。
	array1 为第一个二进制数组， array2 为第二个二进制数组
	时间复杂度为O(n)
*/
func BinaryCompute(a, b []int) []int {
	var result []int
	//进位标志位
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
