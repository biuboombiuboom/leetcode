package algorithm

//RelativeSortArray RelativeSortArray
func RelativeSortArray(arr1 []int, arr2 []int) []int {

	arr := make([]int, 1001)
	l := len(arr1)
	newArr := make([]int, l)

	for i := 0; i < l; i++ {
		arr[arr1[i]] = arr[arr1[i]] + 1
	}

	count := 0
	for _, n := range arr2 {
		for arr[n] > 0 {
			newArr[count] = n
			arr[n] = arr[n] - 1
			count++
		}
	}
	for i, n := range arr {
		for n > 0 {
			newArr[count] = i
			n--
			arr[n] = n
			count++
		}
	}

	return newArr
}
