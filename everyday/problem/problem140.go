package problem

func wordBreak(s string, wordDict []string) []string {
	return wordBreak1(s, toWordMap(wordDict))

}

func wordBreak1(s string, wordMap map[string]struct{}) []string {
	l := len(s)
	sentences := make([]string, 0)

	i, j := 0, 1
	for i < len(s) && j < len(s)+1 {
		word := s[i:j]
		if _, ok := wordMap[word]; ok {
			s = word + " " + s[i+1:]
			i = j
			j++
		} else {
			j++
		}

	}
	if i < len(s) {
		return sentences
	}

	for len(s) > l {

	}

	return sentences
}

func toWordMap(wordDict []string) map[string]struct{} {
	wordMap := make(map[string]struct{})
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}
	return wordMap
}
