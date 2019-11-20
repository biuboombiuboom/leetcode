package problem20

func isValid(s string) bool {
	if s == "" {
		return true
	}

	sm := map[string]string{
		")": "(",
		"]": "{",
		"}": "{",
	}

	var stack []string

	for i := 0; i < len(s); i++ {
		c := s[i : i+1]
		if c1, ok := sm[c]; ok {
			top := ""
			if len(stack) != 0 {
				top = stack[0]
				stack = stack[1:]
			}
			if top != c1 {
				return false
			}
		} else {
			stack = append([]string{c}, stack...)
		}
	}
	return len(stack) == 0
}
