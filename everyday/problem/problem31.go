package problem

func nextPermutation(nums []int) {
	i:=len(nums)-2

	for i>=0&&nums[i]>=nums[i+1]{
		i--
	}
	if i>=0{
		j := len(nums) - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]

	}
	reverse1(nums[i+1:])

	//start:=0
	//end := len(nums) - 1
	//
	//for start<end{
	//	nums[start]=nums[start]+nums[end]
	//	nums[end]=nums[start]-nums[end]
	//	nums[start]=nums[start]-nums[end]
	//	start++
	//	end--
	//}
}

func reverse1(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
