package algorithm

func threeSum(nums []int) [][]int {
	zeroTuple := make([][]int, 0)
	i := 1
	for len(nums)-i > 1 {
		j := i + 1
		num1 := nums[i-1]
		num2 := nums[i]
		for j < len(nums) {
			if num1+num2+nums[j] == 0 {
				zeroTuple = append(zeroTuple, []int{num1, num2, nums[j]})
			}
			j++
		}
		i++
	}

	return zeroTuple
}
