package problem

import (
	"math"
	"sort"
)

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {

	for i := 0; i < len(edges); i++ {
		edges[i] = append(edges[i], i)
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	l := len(edges)

	min_mstw := skip_union(newUnionfind(n), edges, -1)

	keyEdges := make([]int, 0)
	pseudokeyEdges := make([]int, 0)
	for i := 0; i < l; i++ {
		edge := edges[i]
		mst_wight:=skip_union(newUnionfind(n), edges, i)
		if  mst_wight> min_mstw {
			keyEdges = append(keyEdges, edge[3])
			continue
		}
		uf := newUnionfind(n)
		uf.union(edge[0], edge[1])
		if edge[2]+skip_union(uf, edges, i) == min_mstw {
			pseudokeyEdges = append(pseudokeyEdges, edge[3])
		}

	}

	return [][]int{keyEdges, pseudokeyEdges}

}

func skip_union(uf *unionfind, edges [][]int, skip int) int {
	wight_sum := 0
	for i, edge := range edges {
		if i != skip && uf.union(edge[0], edge[1]) {
			wight_sum += edge[2]
		}

	}
	if uf.c > 1 {
		return math.MaxInt64
	}

	return wight_sum
}

func newUnionfind(n int) *unionfind {
	uf := &unionfind{parent: make([]int, n), rank: make([]int, n), c: n}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}

	return uf
}

type unionfind struct {
	parent []int
	rank   []int
	c      int
}

func (uf *unionfind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.parent[fy] = fx
	uf.c--

	return true
}

func (uf *unionfind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}
