package problem20

var sm = make(map[string]string, 3)

func init() {
	sm[")"] = "("
	sm["]"] = "["
	sm["}"] = "{"
}

// //IsValid s
// func IsValid(s string) bool {

// 	if s == "" {
// 		return true
// 	}
// 	if len(s)%2 == 1 {
// 		return false
// 	}

// 	for s != "" {
// 		temp := s
// 		s = strings.Replace(s, "()", "", -1)
// 		s = strings.Replace(s, "[]", "", -1)
// 		s = strings.Replace(s, "{}", "", -1)
// 		if temp == s {
// 			return false
// 		}
// 	}
// 	return true

// }

//IsValid1 s
func IsValid1(s string) bool {
	if s == "" {
		return true
	}

	sm := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var reqEnds []rune
	for _, v := range s {
		char, ok := sm[v]
		if ok {
			reqEnds = append(reqEnds, char)
		} else {
			l := len(reqEnds)
			if l < 1 || reqEnds[l-1] != v {
				return false
			} else {
				reqEnds = reqEnds[:l-1]
			}
		}
	}

	l := len(reqEnds)
	if l > 0 {
		return false
	}
	return true
}
