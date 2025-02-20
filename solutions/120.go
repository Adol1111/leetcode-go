package solutions

func MinimumTotal(triangle [][]int) int {
	// d[i][j] = min(d[i-1][j], d[i-1][j-1]) + triangle[i][j]
	n := len(triangle)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return triangle[0][0]
	}

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, len(triangle[i]))
	}

	var minValue int
	d[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				d[i][j] = d[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				d[i][j] = d[i-1][j-1] + triangle[i][j]
			} else {
				d[i][j] = min(d[i-1][j], d[i-1][j-1]) + triangle[i][j]
			}
			if i == n-1 {
				if j == 0 {
					minValue = d[i][j]
				} else {
					minValue = min(minValue, d[i][j])
				}
			}
		}
	}

	return minValue
}
