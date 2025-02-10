package solutions

func CanCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if gas[0] >= cost[0] {
			return 0
		}
		return -1
	}

	// i的位置要变为 i+cnt+1，因为i->i+cnt已经验证过了
	for i := 0; i < length; i++ {
		totalGas := 0
		totalCost := 0
		cnt := 0
		// 从i开始走，走一圈
		for j := i; cnt < length; j++ {
			// 如果j超过了数组长度，就从头开始
			if j >= length {
				j = 0
			}
			totalGas += gas[j]
			totalCost += cost[j]
			// 无法到达下一个加油站，从新的位置开始，重新走一圈
			if totalGas-totalCost < 0 {
				break
			}
			cnt++
		}
		if cnt == length {
			return i
		}
		// 由于从i->i+cnt 都已验证必定能达到目标，只是剩余油量会有不同，且油量一定是<=当前位置的，所以直接跳过这些位置
		// 因为能到达下一跳，剩余油量一定为正数，从i+1 -> i+cnt 的油量一定会比从i开始的油量少
		i = i + cnt
	}
	return -1
}
