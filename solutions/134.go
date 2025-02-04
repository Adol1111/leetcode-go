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

	for i := 0; i < length; i++ {
		totalGas := 0
		totalCost := 0
		cnt := 0
		for j := i; cnt < length; j++ {
			if j >= length {
				j = 0
			}
			totalGas += gas[j]
			totalCost += cost[j]
			if totalGas-totalCost < 0 {
				break
			}
			cnt++
		}
		if cnt == length {
			return i
		}
		i = i + cnt
	}
	return -1
}
