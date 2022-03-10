package algorithm

import "fmt"

func Print() {
	//grid := [][]byte{
	//	{1,1,0,0,0},
	//	{1,1,0,0,0},
	//	{0,0,1,0,0},
	//	{0,0,0,1,1},
	//}

	board := [][]byte{
		{'O','X','X','O','X'},
		{'X','O','O','X','O'},
		{'X','O','X','O','X'},
		{'O','X','O','O','O'},
		{'X','X','O','X','O'},
	}
	fmt.Println("origin:")
	for _, o := range board {
		fmt.Printf("%s\n",o)

	}
	solve(board)
	fmt.Println("new:")
	for _, o := range board {
		fmt.Printf("%s\n",o)

	}
}
