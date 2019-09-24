package easy

// removeElement 注意边界条件
func removeElement(nums []int, val int) int {
	numsLen := len(nums)
	if numsLen < 1 {
		return 0
	}

	i, j := 0, numsLen - 1
	for i <= j {
		// 这里是关键，如果相等，交换，但是不自增i
		// 如果最后一位和val相等，下次还会检查i位置
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}

	return j + 1
}