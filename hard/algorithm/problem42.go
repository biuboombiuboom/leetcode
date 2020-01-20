package algorithm

func Trap(height []int) int {
	capacity := 0
	maxIndex := 0

	lh := len(height)
	for i := 0; i < lh; i++ {
		if height[i] > height[maxIndex] {
			maxIndex = i
		}
	}
	if lh < 3 {
		return capacity
	}
	l := height[0]
	for i := 1; i < maxIndex; i++ {
		if height[i] < l {
			capacity += l - height[i]
		} else {
			l = height[i]
		}
	}
	l = height[lh-1]
	for i := lh - 2; i > maxIndex; i-- {
		if height[i] < l {
			capacity += l - height[i]
		} else {
			l = height[i]
		}
	}

	return capacity
}
