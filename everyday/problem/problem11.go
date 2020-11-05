package problem

func maxArea(height []int) int {
	maxArea := 0

	l := 0
	r := len(height)-1

	for l<r {
		area:=minI(height[l],height[r])*(r-l)
		if area>maxArea{
			maxArea=area
		}
		if height[l]<height[r]{
			l++
		}else {
			r--
		}
	}

	return maxArea
}


func minI(a, b int) int {
	if a > b {
		return b
	}
	return a
}
