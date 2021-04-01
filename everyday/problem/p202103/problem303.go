package p202103
//
//type NumArray struct {
//	sumMap map[int]int
//}
//
//func Constructor(nums []int) NumArray {
//	sumMap := make(map[int]int)
//	sumMap[-1] = 0
//	for i := 0; i < len(nums); i++ {
//		sumMap[i] += sumMap[i-1] + nums[i]
//	}
//	return NumArray{
//		sumMap,
//	}
//
//}
//
//func (this *NumArray) SumRange(i int, j int) int {
//	return this.sumMap[j] - this.sumMap[i-1]
//}

//给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。
//
//实现 NumArray 类：
//
//NumArray(int[] nums) 使用数组 nums 初始化对象
//int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）
// 
//
//示例：
//
//输入：
//["NumArray", "sumRange", "sumRange", "sumRange"]
//[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
//输出：
//[null, 1, -1, -3]
//
//解释：
//NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
//numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
//numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
//numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
