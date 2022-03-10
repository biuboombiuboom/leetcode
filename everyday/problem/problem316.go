package problem

func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, c := range s {
		left[c-'a']++
	}
	var stack []byte
	inStack := [26]bool{}

	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false

			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}

	return string(stack)
}
