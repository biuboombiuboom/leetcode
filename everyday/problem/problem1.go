package problem

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		num1:=nums[i]
		for j := i+1; j <len(nums) ; j++ {
			num2:=nums[j]
			if num1+num2==target{
				return []int{i,j}
			}
		}
	}
	return nil
}