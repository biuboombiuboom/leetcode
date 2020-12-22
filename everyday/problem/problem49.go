package problem

func groupAnagrams(strs []string) [][]string {
	anagrams := make([][]string, 0)
	anaIndex := make(map[string]int)
	for _, str := range strs {
		newStr := reStore(str)
		if index, ok := anaIndex[newStr]; !ok {
			anaIndex[newStr] = len(anagrams)
			anagrams = append(anagrams, []string{str})

			continue
		} else {
			anagrams[index] = append(anagrams[index], str)
		}
	}

	return anagrams
}

func reStore(str string) string {
	newStr := ""
	cm := make(map[uint8]int)
	for i, _ := range str {
		cm[str[i]]++
	}
	var i uint8
	for i = 'a'; i < 'z'+1; i++ {

		for cm[i] > 0 {
			cm[i]--
			newStr += string(i)
		}
	}
	return newStr
}
