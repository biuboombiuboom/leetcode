package algorithm

func reverseBits(num uint32) uint32 {
	ans := uint32(0)
	n := 31
	for num > 0 {
		ans += (num%2) << n
		num = num >> 1
		n--
	}
	return ans
}
