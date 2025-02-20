package solutions

import "math"

func MySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}

func MySqrt2(x int) int {
	// y = f(x) = x^2 - C
	// 由泰勒展开可得 f(x) = f(x0) + f'(x0)(x-x0) + f''(x0)(x-x0)^2/2! + ... + f(n阶导)(x0)(x-x0)^n/n! + Rn(x)
	// 牛顿迭代法, 即泰勒展开取线性部分(前两项)，并令 f(x) = 0 进行迭代, 因此有 f(x0) + f'(x0)(x-x0) = 0, 即
	// x = x0-f(x0)/f'(x0) , 带入得 x = x0 - (x0^2-C)/(2*x0) = 1/2*(x0 + C/x0)
	// 对以上方程进行多次迭代，可得到比较精确的结果, 当前后两项的差值小于一定的阈值时，即可认为迭代结果是准确的，这里取 1e-7
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		x1 := 0.5 * (C/x0 + x0)
		if math.Abs(x0-x1) < 1e-7 {
			break
		}
		x0 = x1
	}
	return int(x0)
}
