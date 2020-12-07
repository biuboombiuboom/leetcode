package problem

func searchRange1(nums []int, target int) []int {
	ans := []int{-1, -1}
	for i, num := range nums {
		if num == target {
			if ans[0] == -1 {
				ans[0] = i
			}
			ans[1] = i
		}
		if num > target {
			break
		}
	}
	return ans
}

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	l := len(nums)
	if l == 0 {
		return ans
	}
	if l == 1 {
		if nums[0] == target {
			ans[0], ans[1] = 0, 0
		}
		return ans
	}
	mid := nums[l/2]
	if mid == target {
		i := l / 2
		j := l / 2
		for i >= 0 && target == nums[i] {
			i--
		}
		for j < l && target == nums[j] {
			j++
		}

		ans[0] = i + 1
		ans[1] = j - 1
	} else if mid > target {
		ans = searchRange(nums[:l/2], target)
	} else {
		ans = searchRange(nums[l/2:], target)
		if ans[0] > -1 {
			ans[0], ans[1] = ans[0]+l/2, ans[1]+l/2
		}
	}

	return ans

}
