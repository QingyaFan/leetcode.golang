package medium

import (
	"fmt"
	"testing"
)

func TestFindLeftBoundry(t *testing.T) {
	test := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(findLeftBoundry(test, 8))
	fmt.Println(findRightBoundry(test, 8))
}
