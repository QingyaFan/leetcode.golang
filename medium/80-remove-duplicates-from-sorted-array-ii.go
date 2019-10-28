package medium

func removeDuplicates(nums []int) int {
	numsLen := len(nums)

	if numsLen <= 1 {
		return numsLen
	}

	current := 1
	for i := 2; i < numsLen; i++ {
		// i不能加入到current的情况是：i和current还有current-1处的值都相等
		// 能加入的情况是: i和current的值相等&&i和current-1的值不相等、
		//	i和current还有current-1的值都不相等
		//	还有一种情况是不可能出现的情况 ：i和current的值不相等&&i和current的值相等
		// 这种情况是不可能出现的 所以可以用i和current-1的值是否相等来判定是否写入的操作 或者可以用(nums[i]==nums[current]&&nums[i]==nums[current-1])
		if nums[i] != nums[current-1] {
			current++
			nums[current] = nums[i]
		}
	}

	return current + 1
}
