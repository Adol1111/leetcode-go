package solutions

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m, m2 [256]byte

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 0 {
			m[s[i]] = t[i]
		} else if m[s[i]] != t[i] {
			return false
		}
		if m2[t[i]] == 0 {
			m2[t[i]] = s[i]
		} else if m2[t[i]] != s[i] {
			return false
		}
	}
	return true
}
