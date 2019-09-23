package main

// searchInsertPosition common implementation O(n)
func searchInsertPosition(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}

	return numsLen
}

// searchInsertPositionBS implementation using binary search
// O(log(n))
func searchInsertPositionBS(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}