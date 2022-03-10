package algorithm

func numIslands(grid [][]byte) int {

	h := len(grid)
	w := len(grid[0])
	uf := &unionfind{
		parent: make([]int, h*w),
		rank:   make([]int, h*w),
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if int(grid[i][j]) != 0 {
				uf.parent[i*w+j] = i*w + j
				uf.count++
			}
		}
	}

	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {
			if grid[x][y] == 1 {
				k := x*w + y
				grid[x][y] = 0
				if x-1 >= 0 && grid[x-1][y] == 1 {
					uf.union(k, x-1*w+y)
				}
				if x+1 < h && grid[x+1][y] == 1 {
					uf.union(k, x+1*w+y)
				}
				if y-1 >= 0 && grid[x][y-1] == 1 {
					uf.union(k, x*w+y-1)
				}
				if y+1 < w && grid[x][y+1] == 1 {
					uf.union(k, x*w+y+1)
				}
			}
		}
	}

	return uf.count
}

type land struct {
	x int
	y int
}

type unionfind struct {
	parent []int
	rank   []int
	count  int
}

func (uf *unionfind) union(x, y int) {
	fx,fy:=uf.find(x),uf.find(y)
	if fx==fy{
		return
	}
	uf.parent[fy]=fx
	uf.count--
}

func (uf *unionfind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}
