package problem1

//TwoSum t
func TwoSum(nums []int, target int) []int {
	result := make([]int, 2)
	len := len(nums)
	for i := 0; i < len-1; i++ {
		num1 := nums[i]
		num2 := target - num1
		for j := i + 1; j < len; j++ {
			if nums[j] == num2 {
				result[0] = i
				result[1] = j
				break
			}
		}
	}
	return result
}
