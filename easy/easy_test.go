package easy

import (
	"fmt"
	"testing"
)

func TestSearchInsertPosition(t *testing.T) {
	nums, target := []int{1}, 0
	fmt.Println(searchInsertPosition(nums, target))
}

func BenchmarkSearchInsertPosition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums, target := []int{1, 2, 3, 3, 5, 6, 7, 10, 25}, 25
		searchInsertPosition(nums, target)
	}
}

func BenchmarkSearchInsertPositionBS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums, target := []int{1, 2, 3, 3, 5, 6, 7, 10, 25}, 25
		searchInsertPositionBS(nums, target)
	}
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{1, -12, 4, 6}
	fmt.Println(maxSubArray(nums))
}

func TestDeleteNode(t *testing.T) {
	list := []ListNode{}
	node3 := ListNode{3, nil}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	append(list, ListNode{0, &node1})

	deleteNode(&node2)

	// Print the result TODO:
}
