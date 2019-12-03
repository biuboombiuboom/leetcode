package problem38

import "strconv"

import "strings"

import "fmt"

// 1. 1                     1
// 2. 11 1*1                1
// 3. 21 2*1                11
// 4. 1211 1*2+1*1          21
// 5. 111221 1*1+1*2+2*1    1211
// 6. 312211                111221
// 7. 13112221              312211
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := CountAndSay(n - 1)
	fmt.Println(s)
	ch := s[0:1]
	count := 0
	result := make([]string, 2)
	for i := 0; i < len(s); i++ {
		if ch == s[i:i+1] {
			count++
		} else {
			result = append(result, strconv.Itoa(count), ch)
			ch = s[i : i+1]
			count = 1
		}
	}
	result = append(result, strconv.Itoa(count), ch)

	return strings.Join(result, "")
}
