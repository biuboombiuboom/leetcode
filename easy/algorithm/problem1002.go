package algorithm

//CommonChars c
func CommonChars(a []string) []string {
	mm, minIndex := convToMapS(a)
	result := make([]string, 0)
	for i := 0; i < len(a[minIndex]); i++ {
		flag := true
		a1 := a[minIndex][i : i+1]
		for _, m := range mm {
			if mv, ok := m[a1]; !ok {
				flag = false
				break
			} else {
				mv--
				m[a1] = mv
				if mv == 0 {
					delete(m, a1)
				}
			}
		}
		if flag {
			result = append(result, a1)
		}

	}

	return result
}

func convToMapS(ss []string) ([]map[string]int, int) {
	ma := make([]map[string]int, len(ss))
	minLenIndex := 0
	minLen := len(ss[0])
	for i := 0; i < len(ss); i++ {
		s := ss[i]
		l := len(s)
		if l < minLen {
			minLen = l
			minLenIndex = i
		}
		m := make(map[string]int)
		for j := 0; j < l; j++ {
			v := s[j : j+1]
			if mv, ok := m[v]; ok {
				mv++
				m[v] = mv
			} else {
				m[v] = 1
			}
		}

		ma[i] = m
	}
	return ma, minLenIndex
}
