package problem

func reorganizeString(S string) string {
	l := len(S)
	multiParts := make([]string, 0)
	for i := 1; i < len(S); i++ {
		if S[i-1:i] == S[i:i+1] {
			multiParts = append(multiParts, S[i-1:i])
			S = S[:i] + S[i+1:]
			i--
		}
	}
	if len(multiParts) == l {
		return ""
	}
	for i := 0; i < len(multiParts); i++ {
		s := multiParts[i]
		if s != S[0:1] {
			S = s + S
			multiParts= append(multiParts[:i], multiParts[i+1:]...)
			i=-1
			continue
		}
		if s != S[len(S)-1:] {
			S = S + s
			multiParts= append(multiParts[:i], multiParts[i+1:]...)
			i=-1
			continue
		}
		for j := 1; j < len(S)-1; j++ {
			if S[j-1:j] != s && s != S[j:j+1] {
				S = S[:j] + s + S[j:]
				multiParts= append(multiParts[:i], multiParts[i+1:]...)
				i=-1
				break
			}
		}
	}
	if len(multiParts)>0{
		return ""
	}

	return S
}
