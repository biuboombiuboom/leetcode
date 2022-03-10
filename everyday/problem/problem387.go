package problem

func firstUniqChar(s string) int {
	cmap := make(map[int32]int)
	for _, c := range s {
		cmap[c]++
	}
	for i, c := range s {
		if cmap[c] == 1 {
			return i
		}
	}
	return -1
}
