package algorithm

import "strings"

//Convert c
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	sm := make([]string, numRows)

	for i := 0; i < len(s); i++ {
		s1 := s[i : i+1]
		row := 0
		v := (i / (numRows - 1)) % 2
		if v == 0 {
			row = (i) % (numRows - 1)
		} else {
			row = numRows - (i)%(numRows-1) - 1
		}
		sm[row] = sm[row] + s1
	}
	return strings.Join(sm, "")
}
