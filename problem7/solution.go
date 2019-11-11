package problem7

import (
	"strconv"
	"strings"
)

func Reverse(x int) int {
	sx := strconv.Itoa(x)
	len := len(sx)
	result := make([]string, len)
	i := 0
	if sx[0:1] == "-" {
		result[0] = "-"
		i++
		len++
	}
	for ; i < len/2+1; i++ {
		m := sx[i : i+1]
		n := sx[len-i-1 : len-i]
		result[i] = n
		result[len-i-1] = m
	}
	xx, _ := strconv.Atoi(strings.Join(result, ""))
	return xx
}
