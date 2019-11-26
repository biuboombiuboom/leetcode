package problem27

//RemoveElement r
func RemoveElement(nums []int, val int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		}
		i++
	}
	return n
}

//RemoveElement1 r
func RemoveElement1(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	for i := 0; i < len(nums); i++ {
		item := nums[i]
		if item-val == 0 {
			l++
		}
	}
	if l > 0 {
		c := 0
		for i := 0; i < len(nums); i++ {
			item := nums[i]
			if item-val == 0 {
				for j := len(nums) - c - 1; j > i; j-- {
					item1 := nums[j]
					if item1-val != 0 {
						nums[i] = item1
						nums[j] = item
						c++
						break
					}
				}
			}
		}
	}
	return len(nums) - l
}
