package algorithm

//Intersect s
func Intersect(nums1 []int, nums2 []int) []int {
	var m map[int]int
	var lessNum []int
	result := make([]int, 0)
	l1 := len(nums1)
	l2 := len(nums2)

	if l1 > l2 {
		m = convToMap(nums1)
		lessNum = nums2
	} else {
		m = convToMap(nums2)
		lessNum = nums1
	}
	for _, v := range lessNum {
		if mv, ok := m[v]; ok {
			result = append(result, v)
			mv--
			m[v] = mv
			if mv == 0 {
				delete(m, v)
			}
		}
	}

	return result
}

func convToMap(nums []int) map[int]int {
	m := make(map[int]int)

	for _, v := range nums {
		if mv, ok := m[v]; ok {
			mv++
			m[v] = mv
		} else {
			m[v] = 1
		}

	}
	return m
}
