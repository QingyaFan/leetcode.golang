package main

import (
	"fmt"
)

func main() {

	// 01 test
	// var nums []int = []int{0, 2, 4, 5}
	// var target int = 7
	// fmt.Println(TwoSum(nums, target))

	// 02 test
	// var l1 *ListNode = &ListNode{2, nil}
	// l1.Next = &ListNode{4, nil}
	// l1.Next.Next = &ListNode{3, nil}
	// var l2 *ListNode = &ListNode{5, nil}
	// l2.Next = &ListNode{6, nil}
	// l2.Next.Next = &ListNode{4, nil}
	// res := AddTwoNumber(l1, l2)

	// for res != nil {
	// 	fmt.Println(res.Val)
	// 	res = res.Next
	// }

	// 03 test
	// var test string = "pwwkew"
	// fmt.Println(longestUniqueString(test))

	// 07 test
	// fmt.Println(ReverseInteger(123))

	// 13 test
	// fmt.Println(RomanToInteger)

	// 14 test
	// var strs []string = []string{"flower", "flow", "flight"}
	// fmt.Println(LongestCommonPrefix(strs))

	// 25 test
	// fmt.Println(RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	// 70 test
	// fmt.Println(ClimbingStairs(4))

	// 88 test
	// a, b := []int{2,0}, []int{1}
	// MergeSortedArray(a, 1, b, 1)

	// 136 single number
	// nums := []int{4,1,2,1,2}
	// test := SingleNumber2(nums)
	// fmt.Println(test)

	// 189 rotate array
	// a := []int{0, 2, 4, 5}
	// RotateArray(a, 2)
	// fmt.Println(a)

	// 217 contains duplicate
	// a := []int{1,2,3,1}
	// fmt.Println(containsDuplicate(a))

	// 218 contains nearby duplicate
	a := []int{99,99}
	k := 2
	fmt.Println(containsNearbyDuplicate(a, k))

	// 344 reverse string TODO:
	// a := []byte{'h','e','l','l','o'}
	// ReverseString(a)
	// fmt.Printf("%s\n", a)

}
