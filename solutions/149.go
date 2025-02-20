package solutions

func MaxPoints(points [][]int) int {
	// 斜率 = Δx / Δy，由于小数存在精度丢失，所以用最小分数来保存，即
	// 斜率 = Δx / gcd(Δx, Δy)
	// 为方便计算，我们假设分子非负，如果是负数，则分子和分母同时取负即可
	// 由于可能出现垂直的情况，所以斜率不存在，当mx=0时，my=1，当my=0时，mx=1
	// 最终我们得到 斜率 = mx / my ， 用二元组进行保存(mx, my)

	// 由于可能出现重复的点，所以我们用一个map来保存斜率的个数
	// 用一个变量来保存重复的点的个数
	// 用一个变量来保存最大的个数
	// 遍历所有的点，对于每个点，我们遍历所有的点，计算斜率，然后保存到map中
	// 最后遍历map，找到最大的个数

	type slope struct {
		mx int
		my int
	}

	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}

	result := 0
	for i := 0; i < len(points); i++ {
		// result如果已经超过半数或者超过n-i, 则不需要再遍历了, 因为比较的点的个数会越来越少
		if result >= len(points)-i || result > len(points)/2 {
			break
		}
		// 每次遍历都需要清空map，因为可能存在重复的点
		m := make(map[slope]int)
		for j := i + 1; j < len(points); j++ {
			mx := points[j][0] - points[i][0]
			my := points[j][1] - points[i][1]
			if mx == 0 {
				my = 1
			} else if my == 0 {
				mx = 1
			} else {
				if my < 0 {
					mx = -mx
					my = -my
				}
				gcd := gcd(mx, my)
				mx /= gcd
				my /= gcd
				if mx < 0 {
					mx = -mx
					my = -my
				}
			}
			m[slope{mx, my}]++
		}
		for _, v := range m {
			result = max(result, v)
		}
	}
	return result + 1
}
