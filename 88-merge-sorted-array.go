package main

import "fmt"

// MergeSortedArray merge two sorted int array to one
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	if m < 1 {
		for idx, val := range nums2 {
			nums1[idx] = val
		}
	}

	a, b, p := m - 1, n - 1, m + n -1

	for a >= 0 && b >= 0 {

		if b == 0 {
			break
		}

		if a == 0 {
			nums1[b] = nums2[b]
			b--
			continue
		}

		if nums1[a] >= nums2[b] {
			nums1[p] = nums1[a]
			a--
		} else if nums1[a] < nums2[b] {
			nums1[p] = nums2[b]
			b--
		}
		p--
	}

	fmt.Println(nums1)

}
