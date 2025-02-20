package solutions

func AddBinary(a string, b string) string {
	m, n := len(a), len(b)

	l := max(m, n)
	var result = make([]byte, l+1)
	// 第l位会参与 r += int(result[k] - '0') 运算，所以要先补上'0'
	result[l] = '0'

	for i, j, k := m-1, n-1, l; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		var r int
		if i >= 0 {
			r += int(a[i] - '0')
		}
		if j >= 0 {
			r += int(b[j] - '0')
		}
		r += int(result[k] - '0')
		if k > 0 {
			result[k-1] = byte(r/2) + '0'
		}
		result[k] = byte(r%2) + '0'
	}

	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)
}
