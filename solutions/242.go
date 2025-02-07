package solutions

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr [26]int
	for _, v := range s {
		arr[v-'a']++
	}
	for _, v := range t {
		arr[v-'a']--
		if arr[v-'a'] < 0 {
			return false
		}
	}
	return true
}
