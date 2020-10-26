package problem

func smallerNumbersThanCurrent(nums []int) []int {
	l := len(nums)
	numCounts := make([]int, l)
	for i := 0; i < l; i++ {
		num := nums[i]
		sc := 0
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			if num > nums[j] {
				sc++
			}
		}
		numCounts[i] = sc
	}
	return numCounts
}

func smallerNumbersThanCurrent1(nums []int) []int {
	counter := make([]int, 101)

	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	for i := 1; i <= 100; i++ {
		counter[i] += counter[i-1]
	}

	numCounts := make([]int, len(nums))
	for i,v:= range nums {
		if v>0{
			numCounts[i]=counter[v-1]
		}

	}
	return numCounts
}
