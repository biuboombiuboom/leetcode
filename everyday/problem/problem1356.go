package problem

func sortByBits(arr []int) []int {
	sortMap := make(map[int][]int)
	for _, a := range arr {
		count := countOneBit(a)
		if subArr, ok := sortMap[count]; !ok {
			sortMap[count] = []int{a}
		} else {
			inserted := false
			for i := 0; i < len(subArr); i++ {
				if a <= subArr[i] {
					right := append([]int{a}, subArr[i:]...)
					sortMap[count] = append(subArr[0:i], right...)
					inserted = true
					break
				}
			}
			if !inserted {
				sortMap[count] = append(subArr, a)
			}

		}
	}
	ans := make([]int, 0)

	for i := 0; len(sortMap) > 0; i++ {
		ans = append(ans, sortMap[i]...)
		delete(sortMap,i)
	}
	return ans
}

func countOneBit(num int) int {
	oneBitCount := 0
	for num > 0 {
		if num%2 == 1 {
			oneBitCount++
		}
		num = num / 2
	}
	return oneBitCount
}
