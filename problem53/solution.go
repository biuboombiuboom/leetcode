package problem53

//MaxSubArray 1
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := nums[0]
	max := sum
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
