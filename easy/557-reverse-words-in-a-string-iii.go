package easy

import (
	"strings"
)

func reverseWords(s string) string {
	sArr := strings.Split(s, " ")
	sArrLen := len(sArr)
	for i := 0; i < sArrLen; i++ {
		sArr[i] = reverseWord(sArr[i])
	}

	return strings.Join(sArr, " ")
}

func reverseWord(w string) string {
	runes := []rune(w)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
