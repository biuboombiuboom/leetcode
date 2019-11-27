package problem28

//StrStr Sunday
func StrStr(haystack string, needle string) int {
	nlen := len(needle)
	indexM := make(map[string]int)
	for i := 0; i < nlen; i++ {
		s := needle[i : i+1]
		if _, ok := indexM[s]; !ok {
			indexM[s] = i
		}
	}
	for i := 0; i < len(haystack)-nlen; {
		str := haystack[i : i+nlen]
		if str == needle {
			return i
		}
		s1 := haystack[i+1 : i+2]
		if o, ok := indexM[s1]; ok {
			i = i + nlen - o + 1
		} else {
			i = i + nlen + 1
		}
	}
	return -1
}

// func StrStr(haystack string, needle string) int {
// 	i, j := 0, 0
// 	strings.Index(haystack, needle)
// 	for i < len(haystack) && j < len(needle) {
// 		s1 := haystack[i : i+1]
// 		s2 := needle[j : j+1]
// 		if s1 == s2 {
// 			i++
// 			j++
// 		} else {
// 			i = i - j + 1
// 			j = 0
// 		}
// 	}

// 	if j == len(needle) {
// 		return i - j
// 	}
// 	return -1
// }
