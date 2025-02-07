package solutions

func IsHappy(n int) bool {
	if n == 1 {
		return true
	}
	sum := 0
	m := make(map[int]bool)
	for {
		i := n % 10
		sum += i * i
		n /= 10
		if n == 0 {
			if sum == 1 {
				return true
			} else if _, ok := m[sum]; ok {
				return false
			} else {
				n = sum
				m[sum] = true
				sum = 0
			}
		}
	}
}
