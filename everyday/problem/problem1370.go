package problem

func sortString(s string) string {
	bucket := make(map[uint8]int)
	min := uint8(98 + 26)
	max := uint8(98)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := bucket[c]; !ok {
			bucket[c] = 0
		}
		bucket[c] = bucket[c] + 1
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}
	up := true
	cur := min
	ans := ""

	step := 0
	l:=len(bucket)
	for len(bucket) > 0 {
		if _, ok := bucket[cur]; ok {
			ans = ans + string(cur)
			bucket[cur] = bucket[cur] - 1

			if bucket[cur] == 0 {
				delete(bucket, cur)
			}
			step++
			if step == l {
				if up {
					cur = max
				} else {
					cur = min
				}
				up = !up
				step = 0
				l=len(bucket)
			}else{
				if up {
					cur++
				} else {
					cur--
				}
			}
		}		else {
			if up {
				cur++
			} else {
				cur--
			}
		}
	}
	return ans
}

func minStr(s string) string {
	min := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}

	return string(min)
}

func maxStr(s string) string {
	max := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}

	return string(max)
}

var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
