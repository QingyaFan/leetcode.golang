package main

// SingleNumber find the number 
// that appear only once in the array
// TODO:
func SingleNumber(nums []int) int {
	length := len(nums)
	hashtable := make(map[int]int)
	for i := 0; i < length; i++ {
		hashtable[nums[i]] += 1
	}

	for key, val := range hashtable {
		if val == 1 {
			return key
		}
	}

	return -1
}

func SingleNumber2(nums []int) int {
	a := 0
	for _, val := range nums {
		a ^= val
	}

	return a
}
