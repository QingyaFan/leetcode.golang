package easy

// intersection get intersection of two array
// 利用hashmap的key唯一性
// 时间复杂度 O(m+n+k)，k为公共元素个数，所以最坏情况下是O(m+n+min(m,n))
// 空间复杂度 O(k)，k为公共元素，所以最坏情况下是O(min(m,n))
func intersection(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return []int{}
	}

	hash := make(map[int]int)
	var res []int
	for i := 0; i < m; i++ {
		hash[nums1[i]] = 1
	}
	for i := 0; i < n; i++ {
		if hash[nums2[i]] == 1 {
			hash[nums2[i]] = 2
		}
	}

	for key, val := range hash {
		if val > 1 {
			res = append(res, key)
		}
	}

	return res
}
