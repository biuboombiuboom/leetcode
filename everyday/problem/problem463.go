package problem

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		gn := grid[i]
		for j := 0; j < len(gn); j++ {
			if gn[j] == 1 {
				if j-1 < 0 || gn[j-1] == 0 {
					perimeter++
				}
				if j+1 == len(gn) || gn[j+1] == 0 {
					perimeter++
				}
				if i-1 < 0 || grid[i-1][j] == 0 {
					perimeter++
				}
				if i == len(grid)-1 || grid[i+1][j] == 0 {
					perimeter++
				}
			}
		}
	}

	return perimeter
}
