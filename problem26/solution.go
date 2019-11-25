package problem26

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre := nums[0]
	duplicateCount := 0
	l := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == pre {
			duplicateCount++

		} else {
			nums[i-duplicateCount+1] = nums[i]
			pre = nums[i-duplicateCount+1]
			l++
			// duplicateCount-
		}
	}
	return l
}
