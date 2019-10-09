package easy

func reverseStr(s string, k int) string {
	sArr := []rune(s)
	sLen := len(s)
	i := 0
	// i 的递增间隔设置为 2*k 是关键
	for i < sLen {
		if (sLen - i) >= k {
			reverse(sArr[i : i+k])
		} else {
			reverse(sArr[i:])
		}
		i += 2 * k
	}

	return string(sArr)
}

func reverse(s []rune) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
