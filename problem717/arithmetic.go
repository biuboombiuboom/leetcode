package problem717

var (
	chars = []int{0, 10, 11}
)

//IsOneBitCharacter i
func IsOneBitCharacter(bits []int) bool {

	// len := len(bits)
	// zeroIndex := 0
	// for i := 0; i < len; i++ {
	// 	char := bits[i]
	// 	if char == 1 {
	// 		i++
	// 	}
	// 	if char == 0 {
	// 		zeroIndex = i
	// 	}
	// }
	// return zeroIndex == len-1
	len := len(bits)
	i := 0
	for i < len-1 {
		i += bits[i] + 1
	}
	return i == len-1
}
