package solutions

func SingleNumber2(nums []int) int {
	// 统计每个数字出现的次数
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	// 找出出现一次的数字
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return 0
}

func SingleNumber22(nums []int) int {
	// 原理是设计一个电路，让3个相同数字转为0，即 00 -> 01 -> 10 -> 00

	// | (ai,bi) |xi | 新的(ai,bi) |
	// | -- | -- | -- |
	// | 00 | 0 | 00 |
	// | 00 | 1 | 01 |
	// | 01 | 0 | 01 |
	// | 01 | 1 | 10 |
	// | 10 | 0 | 10 |
	// | 10 | 1 | 00 |

	// 如果只考虑输出ai
	// | (ai,bi) |xi | 新的ai |
	// | -- | -- | -- |
	// | 00 | 0 | 0 |
	// | 00 | 1 | 0 |
	// | 01 | 0 | 0 |
	// | 01 | 1 | 1 |
	// | 10 | 0 | 1 |
	// | 10 | 1 | 0 |

	// 将bi换成新的bi
	// | (ai,新的bi) |xi | 新的ai |
	// | -- | -- | -- |
	// | 00 | 1 | 0 |
	// | 01 | 0 | 0 |
	// | 01 | 1 | 1 |
	// | 00 | 0 | 1 |
	// | 10 | 1 | 0 |
	// | 10 | 0 | 0 |

	// 即ai=ai′bi′xi+aibi′xi′=bi′(ai⊕xi) // 需要电路知识
	// 因此可以得到2个表达式
	// 1. b = ~a & (bˆx)
	// 2. a = ~b & (aˆx)

	a, b := 0, 0
	for _, num := range nums {
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}
