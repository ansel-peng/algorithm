package section1

/*
	ReversePair
	逆序对
 	nums 数组
*/

func ReversePairs(nums []int) int {
	return mergeSortForReversePairs(nums, 0, len(nums)-1)
}

func mergeSortForReversePairs(nums []int, start, end int) int {
	var tmp []int
	if start >= end {
		return 0
	}
	mid := start + (end - start)/2
	cnt := mergeSortForReversePairs(nums, start, mid) + mergeSortForReversePairs(nums, mid + 1, end)
	i, j := start, mid + 1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i - start]
	}
	return cnt

}
