package algorithm

func threeSum(nums []int) [][]int {
	zeroTuple := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			for k := 0; k < len(nums); k++ {
				if k == i || k == j {
					continue
				}
				num1 := nums[i]
				num2 := nums[j]
				num3 := nums[k]

				if num1+num2+num3==0{
				}
			}
		}
	}

	return zeroTuple
}

