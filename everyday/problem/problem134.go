package problem

/*
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
[4,5,2,6,5,3]
[3,2,7,3,2,9]
输出: 3
*/
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		remain := gas[i]
		j, count := i, 0
		for count <= len(gas) {

			remain -= cost[j]
			if remain < 0 {
				if j>i{
					i=j
				}
				break
			}
			j++
			if j >= len(gas) {
				j = 0
			}
			remain += gas[j]
			count++
		}
		if remain >= 0 {
			return i
		}

	}

	return -1
}
