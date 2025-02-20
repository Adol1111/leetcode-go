package solutions

func ReverseBits(num uint32) uint32 {
	var rev uint32
	for i := 0; i < 32 && num > 0; i++ {
		// 二进制右边是低位，左边是高位
		// 先通过与1&取出最低位，然后左移31-i位，填充到rev的第31-i位上
		// 最低位取出后，num右移一位，这样i+1位就会变为最低位
		rev |= (num & 1) << (31 - i)
		num >>= 1
	}
	return rev
}
