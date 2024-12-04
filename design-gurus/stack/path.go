package stack

import "strings"

func SimplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	pathStack := make([]string, 0)

	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		}

		if dir == ".." && len(pathStack) > 0 {
			pathStack = pathStack[:len(pathStack)-1]
		} else if dir != ".." {
			pathStack = append(pathStack, dir)
		}
	}

	return "/" + strings.Join(pathStack, "/")
}
