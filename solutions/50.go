package solutions

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / MyPow(x, -n)
	}

	// x^n = x^(2*n/2) = (x^2)^(n/2)
	// 令 n=2k, x^2k = (x^2)^k
	// 令 n=2k+1, x^(2k+1) = (x^2)^(k+1/2) = x*(x^2)^k
	if n%2 == 0 {
		return MyPow(x*x, n/2)
	}
	return x * MyPow(x*x, n/2)
}
