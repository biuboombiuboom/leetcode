package algorithm

import "strings"

func SimplifyPath(path string) string {
	var stack []string

	paths := strings.Split(path, "/")

	for i := 0; i < len(paths); i++ {
		if paths[i] == "." {
			continue
		}
		if paths[i] != ".." && paths[i] != "" {
			stack = append(stack, paths[i])
			continue
		}
		if paths[i] == ".." && len(stack) > 0 {
			stack = stack[0 : len(stack)-1]
		}

	}
	return "/" + strings.Join(stack, "/")
}
