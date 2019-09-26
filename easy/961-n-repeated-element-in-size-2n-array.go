package easy

// repeatedNTimes 找到重复N次的元素
// 但从题目中描述，2N个元素，有N-1个不同值
// 也就是说只有一个元素重复，所以只要找到
// 某元素重复超过2次，返回就可以了
// 时间复杂度 O(<n)
// 空间复杂度 O(n/2)
func repeatedNTimes(A []int) int {
	hash := make(map[int]int)
	for i := 0; i < len(A); i++ {
		hash[A[i]]++
		if hash[A[i]] > 1 {
			return A[i]
		}
	}
	panic(-1)
}
