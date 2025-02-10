package solutions

func CanConstruct(ransomNote string, magazine string) bool {
	arr := [26]int{}
	// 统计ransomNote的所有字符数
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
