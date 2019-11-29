package problem29

import "math"

func divide(dividend int, divisor int) int {

	result := 0
	if result > math.MaxInt32 || result < math.MinInt32 {
		return math.MaxInt32
	}
	return result

}
