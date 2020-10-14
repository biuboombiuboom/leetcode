package offer

func findRepeatNumber1(nums []int) int {
	buckets := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := buckets[num]; ok {
			return num
		} else {
			buckets[num] = 1
		}
	}
	return 0
}

func findRepeatNumber(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {

	}
}
