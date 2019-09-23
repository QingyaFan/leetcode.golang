package main

import (
	"fmt"
	"testing"
)

func TestSearchInsertPosition(t *testing.T)  {
	nums, target := []int{1}, 0
	fmt.Println(searchInsertPosition(nums, target))
}

func BenchmarkSearchInsertPosition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums, target := []int{1,2,3,3,5,6, 7,10,25}, 25
		searchInsertPosition(nums, target)
	}
}

func BenchmarkSearchInsertPositionBS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums, target := []int{1,2,3,3,5,6, 7,10,25}, 25
		searchInsertPositionBS(nums, target)
	}
}