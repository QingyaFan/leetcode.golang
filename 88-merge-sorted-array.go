package main

// MergeSortedArray merge two sorted int array to one
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	a, b, p := m - 1, n - 1, m + n -1

	for a >= 0 && b >= 0 {
		if nums1[a] >= nums2[b] {
			nums1[p] = nums1[a]
			a--
		} else if nums1[a] < nums2[b] {
			nums1[p] = nums2[b]
			b--
		}
		p--
	}

	for i := 0; i < b + 1; i++ {
		nums1[i] = nums2[i]
	}

}
