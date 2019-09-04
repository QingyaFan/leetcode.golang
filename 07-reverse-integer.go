package main

// ReverseInteger rever integer
func ReverseInteger(x int) int {

	const MIN int = 0x80000000
	const MAX int = 0x7fffffff

	var res int

	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}

	if res > MAX || res < -MIN {
		return 0
	}

	return res
}
