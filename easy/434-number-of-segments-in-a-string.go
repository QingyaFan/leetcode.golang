package easy

func countSegments(s string) int {
	var segments int
	for i := 0; i < len(s); i++ {
		// 在当前字符不为空格的情况下，有两种情况有单词
		// 1、在首个单词，即 i=0
		// 2、前一个字符为空格
		// 相当于对首个单词特殊处理，后面一个单词对应一个空格
		if s[i] != ' ' && (i == 0 || s[i-1] == ' ') {
			segments++
		}
	}

	return segments
}
