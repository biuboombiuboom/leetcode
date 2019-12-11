package algorithm

//SortColors s
func SortColors(nums []int) {
	redRightIndex := 0
	blueLeftIndex := len(nums) - 1
	i := 0
	for i <= blueLeftIndex {
		color := nums[i]
		switch color {
		case 0: //红色
			nums[i] = nums[redRightIndex]
			nums[redRightIndex] = color
			redRightIndex++
			i++
		case 1:
			i++
		case 2: //蓝色
			nums[i] = nums[blueLeftIndex]
			nums[blueLeftIndex] = color
			blueLeftIndex--
		}
	}
}

//SortColors1 桶排
func SortColors1(nums []int) {
	brokers := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		brokers[nums[i]]++
	}
	c := 0
	for i := 0; i < len(brokers); i++ {
		for j := 0; j < brokers[i]; j++ {
			nums[c] = i
			c++
		}
	}
}
