package main

func RemoveDuplicates(nums []int) int {
	var i int = 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i += 1
			nums[i] = nums[j]
		}
	}

	return i + 1
}
