package main

// ReverseString reverse string
func ReverseString(s []byte)  {
	a, b := 0, len(s) - 1
	for a < b {
		s[a], s[b] = s[b], s[a]
		a++; b--
	}
}