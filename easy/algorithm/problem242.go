package algorithm

//IsAnagram i
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := toMap(s)
	for i := 0; i < len(t); i++ {
		t1 := t[i : i+1]
		if c, ok := m[t1]; ok {
			c--
			m[t1] = c
			if c == 0 {
				delete(m, t1)
			}
		} else {
			return false
		}
	}
	return len(m) == 0
}

func toMap(s string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		s1 := s[i : i+1]
		if c, ok := m[s1]; !ok {
			m[s1] = 1
		} else {
			c++
			m[s1] = c
		}
	}
	return m
}
