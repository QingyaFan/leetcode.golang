package easy

import "math"

// maxSubArray
// 动态规划方法，子序列有太多种可能，
// 这正好符合动态规划解决子问题的特征
func maxSubArray(nums []int) int {
	ans, sum := nums[0], 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		ans = int(math.Max(float64(ans), float64(sum)))
	}

	return ans
}
