package problem

import "sort"

func minCostConnectPoints(points [][]int) int {
	edges := make([]edge, 0)
	parent := make([]int, len(points))
	for i, point := range points {
		parent[i] = i
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, edge{i, j, dist(point, points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})
	uf := newUnionFind(len(points))
	left := len(points) - 1
	ans := 0
	for _, e := range edges {
		if uf.union(e.a, e.b) {
			ans += e.dist
			left--
			if left == 0 {
				break
			}
		}
	}
	return ans
}

func dist(a, b []int) int {
	return abs_(a[0]-b[0]) + abs_(a[1]-b[1])
}

func abs_(num int) int {
	if num < 0 {
		num = num * -1
	}
	return num
}

type edge struct {
	a    int
	b    int
	dist int
}

type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i]=i
	}
	return &unionFind{parent, rank}
}

func (uf *unionFind) union(a, b int) bool {
	fx, fy := uf.find(a), uf.find(b)
	if fx == fy {
		return false
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.parent[fy] = fx
	return true

}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}


func run(){
	points:=[][]int{
		{1,2},{4,2},{3,9},{8,15},
	}
	minCostConnectPoints(points)
}