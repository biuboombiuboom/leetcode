package problem

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		l1 := l - i*2
		k:=0
		for j := i; j <l1-1+i; j++ {
			matrix[i][j], matrix[j][l1-1+i], matrix[l1-1+i][l1+i-1-k], matrix[l1+i-1-k][i]= matrix[l1+i-1-k][i],matrix[i][j], matrix[j][l1-1+i], matrix[l1-1+i][l1+i-1-k]
			k++

		}
	}

}
