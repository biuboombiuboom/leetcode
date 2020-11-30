package problem

func moveZeroes(nums []int) {
	l := len(nums)
	count := 0
	for i := 0; i < l&&count<len(nums); {
		if nums[i] == 0 {
			count++
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			continue
		}
		i++
	}
}

func moveZeroes1(nums []int) {
	l := len(nums)
	j := l - 1
	count := 0
	for i := l - 1; i >= 0&&count<len(nums); i-- {
		if nums[i] == 0 {
			if j != -1 {
				nums[i], nums[j] = nums[j], nums[i]
				if j == l-count-1 {
					count++
				}
				i = l - count-1
			}
		}
		j = i
	}

}
