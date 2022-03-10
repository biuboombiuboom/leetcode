package problem

func checkPossibility(nums []int) bool {
	n:=len(nums)
	if n<3{
		return  true
	}

	c:=0
	for i := 0; i <n-1 ; i++ {
		if nums[i]>nums[i+1]{
			c++
			if i>0&&nums[i+1]<nums[i-1]{
				nums[i+1]=nums[i]
			}
		}
		if c>1{
			return  false
		}
	}
	return  true
}
