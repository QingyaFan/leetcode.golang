package easy

// ClimbingStairs climb stairs
func ClimbingStairs(n int) int {

	if n <= 3 {
		return n
	}

	var total, minusTwo, minusOne int = 0, 2, 1

	for i := 3; i <= n; i++ {
		total = minusTwo + minusOne
		minusOne = minusTwo
		minusTwo = total
	}

	return total
}
