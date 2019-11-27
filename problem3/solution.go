package problem3

//LengthOfLongestSubstring sliding window
func LengthOfLongestSubstring(s string) int {
	l := len(s)
	sm := make(map[string]int)
	max := 0
	left := 0
	for i := 0; i < l; i++ {
		s1 := s[i : i+1]
		if o, ok := sm[s1]; ok {
			left = maxx(o+1, left)
		}
		sm[s1] = i
		max = maxx(max, i-left+1)
	}
	return max
}

func maxx(a, b int) int {
	if a-b > 0 {
		return a
	}
	return b
}

//LengthOfLongestSubstring l 暴力破解
// func LengthOfLongestSubstring(s string) int {
// 	l := len(s)
// 	if l < 2 {
// 		return l
// 	}
// 	max := 0
// 	for i := 0; i < l; {
// 		sm := make(map[string]int)
// 		for j := i; j < l; j++ {
// 			s1 := s[j : j+1]
// 			if _, ok := sm[s1]; !ok {
// 				sm[s1] = j
// 			} else {
// 				i = sm[s1] + 1
// 				break
// 			}
// 			if j == l-1 {
// 				if len(sm) > max {
// 					max = len(sm)
// 				}
// 				return max
// 			}
// 		}
// 		if len(sm) > max {
// 			max = len(sm)
// 		}
// 	}
// 	return max
// }
