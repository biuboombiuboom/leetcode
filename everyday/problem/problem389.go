package problem

func findTheDifference(s string, t string) byte {
	diff := 0
	for i := 0; i < len(s); i++ {
		diff += int(t[i] - s[i])
	}
	diff += int(t[len(s)])

	return uint8(diff)
}
