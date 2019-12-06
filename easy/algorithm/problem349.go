package algorithm

//Intersection i
func Intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	result := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		n1 := nums1[i]
		for j := 0; j < len(nums2); j++ {
			n2 := nums2[j]
			if n1 == n2 {
				if _, ok := m[n2]; !ok {
					result = append(result, n2)
					m[n2] = n2
				}
			}
		}
	}

	return result
}
