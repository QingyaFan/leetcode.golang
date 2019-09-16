package main

// containsDuplicate given a int array
// return if there are repeat nums
func containsDuplicate(nums []int) bool  {
	numsLen := len(nums)
	if numsLen <= 1 {
		return false
	}

	hash := make(map[int]int)
	for i := 0; i < numsLen; i++ {
		if hash[nums[i]] > 0 {
			return true
		}
		hash[nums[i]]++
	}

	return false
}