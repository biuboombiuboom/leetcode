package problem

func singleNumbers(nums []int) []int {
	result:=[]int{0,0}
	xorSum:=0
	for i := 0; i <len(nums) ; i++ {
		xorSum^=nums[i]
	}
	lowbit:=xorSum&(-xorSum)
	for i := 0; i <len(nums) ; i++ {
		x:=nums[i]
		if x&lowbit >0{
			result[0]^=x
		}else{
			result[1]^=x
		}
	}
	return result
}