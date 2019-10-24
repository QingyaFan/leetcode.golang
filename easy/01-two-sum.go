package easy

// TwoSum 给定一个整数数组 nums 和一个目标值 target，
// 请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
func TwoSum(nums []int, target int) []int {
	elMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		var el = target - nums[i]
		if _, exist := elMap[el]; exist {
			return []int{elMap[el], i}
		}
		elMap[nums[i]] = i
	}
	return []int{-1, -1}
}
