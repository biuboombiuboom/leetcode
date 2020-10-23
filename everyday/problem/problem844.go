package problem

func backspaceCompare(S string, T string) bool {
	back := "#"
	ls := len(S)
	lt := len(T)

	ss := 0
	st := 0 //退格计数
	i, j := ls-1, lt-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if S[i:i+1] == back {
				ss++
				i--
			} else if ss > 0 {
				ss--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if T[j:j+1] == back {
				st++
				j--
			} else if st > 0 {
				st--
				j--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}

		i--
		j--

	}

	return true

}
