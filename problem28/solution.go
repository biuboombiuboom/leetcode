package problem28

import "strings"

//StrStr s
func StrStr(haystack string, needle string) int {
	i, j := 0, 0
	strings.Index(haystack, needle)
	for i < len(haystack) && j < len(needle) {
		s1 := haystack[i : i+1]
		s2 := needle[j : j+1]
		if s1 == s2 {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == len(needle) {
		return i - j
	}
	return -1
}
