package problem218

func getSkyline(buildings [][]int) [][]int {

	count := len(buildings)
	if count == 0 {
		return buildings
	}
	result := make([][]int, 1)
	result = append(result, []int{buildings[0][0], buildings[0][2]})
	for i := 0; i < count; i++ {

	}
	return result
}
