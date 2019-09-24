package easy

// PalindromeNumber palindrome number
func PalindromeNumber(x int) bool {

	// 负数不可能是回文数
	// 末尾是0，且x不是0的，不可能是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedNum := 0
	for x > reversedNum {
		reversedNum = x%10 + reversedNum*10
		x /= 10
	}

	// 如果x是奇数位数，中间的数字不影响回文，可以直接去掉
	return reversedNum == x || reversedNum/10 == x
}
