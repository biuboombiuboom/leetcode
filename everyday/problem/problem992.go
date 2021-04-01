package problem

func subarraysWithKDistinct(A []int, K int) int {
	n := len(A)
	subarrayLen := 0
	subMap := make(map[int][]int)
	ans := 0
	for i := 0; i < n; i++ {
		a := A[i]
		if len(subMap) < K {
			if _, ok := subMap[a]; !ok {
				subMap[a] = make([]int, 0)
			}
			subMap[a] = append(subMap[a], i)
			if len(subMap)==K{
				ans++
			}
			subarrayLen++
		} else if _, ok := subMap[a]; ok {
			subMap[a] = append(subMap[a], i)
			subarrayLen++
			ans++
		} else {
			for len(subMap)==3 {
				firstSubElem := A[i-subarrayLen]
				if len(subMap[firstSubElem]) == 1 {
					delete(subMap, firstSubElem)
					subarrayLen--
					i=i-subarrayLen
				} else {
					subMap[firstSubElem] = subMap[firstSubElem][1:]
					subarrayLen--
					ans++
				}

			}

		}
	}
	return ans
}
