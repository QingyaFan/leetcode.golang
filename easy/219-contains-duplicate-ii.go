package easy

func containsNearbyDuplicate(nums []int, k int) bool {
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
		if len(hash) > k {
			delete(hash, nums[i-k])
		}
	}

	return false
}