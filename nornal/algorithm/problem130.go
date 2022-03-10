package algorithm

func solve(board [][]byte) {
	h := len(board)
	if h == 0 {
		return
	}
	w := len(board[0])
	uf := newUnionfind(h * w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == 'O' {
				uf.parent[i*w+j] = i*w + j
				uf.rank[i*w+j] = 1
			}
		}
	}
	uf.parent = append(uf.parent, h*w)
	uf.rank = append(uf.rank, h*w)
	o_board := make([][]int, 0)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == 'O' {
				o_board = append(o_board, []int{i, j})
				if i == 0 || j == 0 || i == h-1 || j == w-1 {
					uf.union(i*w+j,h*w )
					continue
				}
				if board[i-1][j] == 'O' {
					uf.union( i*w+j,(i-1)*w+j)
				}
				if board[i+1][j] == 'O' {
					uf.union(i*w+j,(i+1)*w+j )
				}
				if board[i][j-1] == 'O' {
					uf.union(i*w+j,i*w+j-1 )
				}
				if board[i][j+1] == 'O' {
					uf.union(i*w+j,i*w+j+1)
				}
			}
		}
	}

	for i := 0; i < len(o_board); i++ {
		x,y:=o_board[i][0],o_board[i][1]
		if uf.find(h*w) != uf.find(x*w+y) {
			board[x][y] = 'X'
		}
	}

}

func newUnionfind(n int) *unionfind1 {
	parent := make([]int, n)
	rank := make([]int, n)

	return &unionfind1{
		parent, rank,
	}
}

func (uf *unionfind1) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fy] += uf.rank[fx]
	uf.parent[fy] = fx

}

func (uf *unionfind1) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

type unionfind1 struct {
	parent []int
	rank   []int
}
