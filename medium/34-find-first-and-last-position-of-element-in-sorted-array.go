package medium

import "fmt"

// searchRange 给定一个按照升序排列的整数数组 nums，
// 和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置.
// 时间复杂度要求log(n)，不存在返回[-1,-1]
func searchRange(nums []int, target int) []int {
	var res []int

	res = append(res, findLeftBoundry(nums, target))
	res = append(res, findRightBoundry(nums, target))

	return res
}

func findLeftBoundry(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		middle := (left + right) / 2
		if nums[middle] == target {
			right = middle
		}
		if nums[middle] > target {
			right = middle
		}
		if nums[middle] < target {
			left = middle + 1
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func findRightBoundry(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		middle := (left + right + 1) / 2
		if nums[middle] == target {
			left = middle
		}
		if nums[middle] < target {
			left = middle + 1
		}
		if nums[middle] > target {
			right = middle - 1
		}
		fmt.Println(left, middle, right)
	}

	if left > len(nums)-1 {
		return -1
	}
	if nums[left] == target {
		return left
	}

	return -1
}
