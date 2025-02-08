package solutions

func SimplifyPath(path string) string {
	stack := make([]string, 0)
	for i := 0; i < len(path); {
		if path[i] == '/' {
			i++
			continue
		}
		j := i
		for ; j < len(path) && path[j] != '/'; j++ {
		}
		dir := path[i:j]
		i = j
		if dir == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if dir != "." && dir != "" {
			stack = append(stack, dir)
		}
	}
	var result string
	for _, dir := range stack {
		result += "/" + dir
	}
	if result == "" {
		return "/"
	}
	return result
}
