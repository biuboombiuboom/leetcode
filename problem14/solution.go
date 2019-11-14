package problem14

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	p1 := ""
	for i := 0; i < len(strs[0]); i++ {
		p1 = strs[0][0 : i+1]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i {
				return p1[0:i]
			}
			p := strs[j][0 : i+1]
			if p1 != p {
				return p1[0:i]
			}
		}
	}
	return p1
}
