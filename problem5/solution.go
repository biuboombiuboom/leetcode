package problem5

//LongestPalindrome1 中心扩展
func LongestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	result := s[0:1]
	for i := 0; i < len(s); i++ {
		L1, R1 := i, i
		//寻奇数回文
		for L1 > -1 && R1 < len(s) && s[L1] == s[R1] {
			L1--
			R1++
		}
		l := R1 - L1 - 1
		L2, R2 := i, i+1
		//寻偶数回文
		for L2 > -1 && R2 < len(s) && s[L2] == s[R2] {
			L2--
			R2++
		}
		if l < R2-L2-1 {
			l = R2 - L2 - 1
		}

		if len(s[i-l/2:i+l/2]) > len(result) {
			result = s[i-l/2 : i+l/2]
		}

	}
	return result
}

//LongestPalindrome 暴力破解
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	result := s[0:1]
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			s1 := s[i : j+1]

			flag := true
			left := len(s1)/2 - 1
			right := len(s1)/2 + 1
			if len(s1)%2 == 0 {
				right--
			}
			// if len(s1)%2 == 1 {
			// 	left++
			// }
			for left > -1 && right < len(s1) {
				if s1[left] != s1[right] {
					flag = false
					break
				}
				left--
				right++
			}
			if flag && len(s1) > len(result) {
				result = s1
			}
		}
	}

	return result
}
