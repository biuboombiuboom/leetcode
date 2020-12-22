package problem

import "strings"

func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")
	if len(sList) != len(pattern) {
		return false
	}
	pIndex := make(map[uint8]int)
	sIndex := make(map[string]int)
	for i, str := range sList {
		pc:=pattern[i]
		if pIndex[pc]!=sIndex[str]{
			return  false
		}
		pIndex[pc]=sIndex[str]+(i+1)*10
		sIndex[str]=sIndex[str]+(i+1)*10
	}

	return true
}
