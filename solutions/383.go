package solutions

func CanConstruct(ransomNote string, magazine string) bool {
	arr := [26]int{}
	for i := 0; i < len(ransomNote); i++ {
		arr[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
		arr[magazine[i]-'a']--
	}
	for _, v := range arr {
		if v > 0 {
			return false
		}
	}
	return true
}
