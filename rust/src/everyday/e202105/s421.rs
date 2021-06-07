struct Solution;
impl Solution {
    pub fn find_maximum_xor(nums: Vec<i32>) -> i32 {
        let n=nums.len();
        let mut max=0;
        for i in 0..n{
            for j in i+1..n{
                max=i32::max(max, nums[i]^nums[j])
            }
        }
        max
    }
}

pub fn run(){
    println!("{}", Solution::find_maximum_xor(vec![1]));
}