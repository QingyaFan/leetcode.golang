package easy

// intersect 计算两个数组的交集
func intersect(nums1 []int, nums2 []int) []int {
	nums1Len, nums2Len := len(nums1), len(nums2)

	// 边界情况
	if nums1Len == 0 || nums2Len == 0 {
		return []int{}
	}

	hashMap := make(map[int]int)
	for i := 0; i < nums1Len; i++ {
		hashMap[nums1[i]]++
	}

	var res []int
	for i := 0; i < nums2Len; i++ {
		if hashMap[nums2[i]] > 0 {
			res = append(res, nums2[i])
			hashMap[nums2[i]]--
		}
	}

	return res
}

// 灵魂三问：
// 1. 如果给定的数组已经排好序，使用双指针法更高效
// 2. 其中一个数组比另一个数组小很多，使用双指针更高效
// 3. 如果nums2不能放到内存中了，在磁盘上，我们可以把nums1放入hashMap中，
//    以流的方式读nums2，逐一比对
//    如果两个数组都足够大，不能放入到内存中，我们可以都单独排下序，然后用双指针法，
//    每次分别读两个数组的一个元素进行比对
