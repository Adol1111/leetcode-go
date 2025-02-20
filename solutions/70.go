package solutions

func ClimbStairs(n int) int {
	// dn = dn-1 + dn-2
	if n <= 2 {
		return n
	}

	// 递归会超时
	// return ClimbStairs(n-1) + ClimbStairs(n-2)

	f1, f2 := 1, 2
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}
