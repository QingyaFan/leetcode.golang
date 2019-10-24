package medium

func longestUniqueString(str string) int {

	strLen := len(str)
	subStr := make(map[byte]int)
	res := 0

	for i, j := 0, 0; j < strLen; j++ {
		if v, ok := subStr[byte(str[j])]; ok {
			if i < v {
				i = v
			}
		}
		if res < (j - i + 1) {
			res = j - i + 1
		}
		subStr[byte(str[j])] = j + 1
	}

	return res
}
