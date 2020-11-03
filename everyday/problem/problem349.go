package problem

func intersection(nums1 []int, nums2 []int) []int {
	sectionMap := make(map[int][]int)
	for _, num := range nums1 {
		if _, ok := sectionMap[num]; !ok {
			sectionMap[num] = make([]int, 2)
		}
		sectionMap[num][0]++
	}

	for _, num := range nums2 {
		if _, ok := sectionMap[num]; !ok {
			sectionMap[num] = make([]int, 2)
		}
		sectionMap[num][1]++
	}
	section := make([]int, 0)
	for k, v := range sectionMap {
		if v[0] > 0 && v[1] > 0 {
			section = append(section, k)
		}
	}
	return section
}
