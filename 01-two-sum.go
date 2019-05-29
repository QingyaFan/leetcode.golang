package main

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
