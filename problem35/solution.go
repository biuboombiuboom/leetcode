package problem35

//SearchInsert s
func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if nums[right] < target {
		return right + 1
	}
	for left <= right {
		mid := (right + left) / 2
		num := nums[mid]
		if target == num {
			return mid
		}
		if target > num {
			left = mid + 1
		}
		if target < num {
			right = mid - 1
		}
	}
	return left
}

//SearchInsert1 s
func SearchInsert1(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		if nums[0] < target {
			return 1
		}
		return 0
	}
	i := l / 2
	num := nums[i]
	if num == target {
		return i
	}
	if num > target {
		return SearchInsert(nums[0:i], target)
	}
	index := i + SearchInsert(nums[i:], target)

	return index
}
