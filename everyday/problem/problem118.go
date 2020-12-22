package problem

func generate(numRows int) [][]int {
	rows := make([][]int, 0)
	if numRows > 0 {
		for i := 0; i < numRows; i++ {
			row := make([]int, i+1)
			rows = append(rows, row)
			for j := 0; j < i+1; j++ {
				if j == 0 || j == len(row)-1 {
					row[j] = 1
				} else {
					row[j] = rows[i-1][j-1] + rows[i-1][j]
				}
			}
		}
	}

	return rows
}
