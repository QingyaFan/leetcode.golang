package easy

func rotate(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// RotateArray rotate an array
func RotateArray(nums []int, k int) {
	numsLen := len(nums)

	if k == numsLen || numsLen <= 1 {
		return
	}

	n := numsLen - (k % numsLen)
	rotate(nums[:n])
	rotate(nums[n:])
	rotate(nums)
}
