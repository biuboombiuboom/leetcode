package problem

import (
	"strconv"
	"strings"
)

func splitIntoFibonacci(S string) []int {
	max := 1<<31 - 1
	ans := make([]int, 0)
	if len(S) < 3 {
		return ans
	}
	i := 1
	step := 1
	for len(S)-i > i {
		back := false
		j := i + step
		one, _ := strconv.ParseUint(strings.TrimLeft(S[0:i], "0"), 0, 32)
		two, _ := strconv.ParseUint(strings.TrimLeft(S[i:j], "0"), 0, 32)
		ans = append(ans, int(one))
		ans = append(ans, int(two))

		c := j
		for j < len(S) {
			three := int(one + two)
			if three > max {
				ans = make([]int, 0)
				break
			}
			sum := strconv.Itoa(three)
			k := len(sum) + j
			if k > len(S) || S[j:k] != sum {
				step++
				back = true
				break
			}
			j = k
			ans = append(ans, three)
			one = two
			two = uint64(three)
			c += len(sum)
		}

		if c == len(S) && len(ans) > 2 {
			return ans
		}
		ans = make([]int, 0)
		if !back {
			i++
			step=1
		}
	}
	return ans
}
