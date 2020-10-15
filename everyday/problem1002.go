package main

import "sort"

func commonChars(A []string) []string {
	if len(A) == 1 {
		return A
	}
	amap := make(map[string][]int)
	for i := 0; i < len(A[0]); i++ {
		a := string(A[i])
		amap[a] = make([]int, len(A))
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			a := string(A[i][j])
			if _, ok := amap[a]; ok {
				amap[a][i] = amap[a][i] + 1
			}
		}
	}
	w := make([]string, 0)
	for k, v := range amap {
		sort.Ints(v)
		for i := 0; i < v[0]; i++ {
			w = append(w, k)
		}
	}
	return w
}
