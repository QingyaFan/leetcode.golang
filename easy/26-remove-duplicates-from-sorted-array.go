package easy

// RemoveDuplicates remove duplicates from a array
func RemoveDuplicates(nums []int) int {
	var i int
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i ++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
