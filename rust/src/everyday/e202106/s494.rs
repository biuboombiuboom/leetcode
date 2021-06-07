
struct Solution;
impl Solution {
    pub fn find_target_sum_ways(nums: Vec<i32>, target: i32) -> i32 {
        let mut dp=Vec::new();
    
        let n=nums.len();
        let mut sum=0;
        for i in 0..n{
            sum+=nums[i];            
        }
        let diff=sum-target;

        if diff<0|| diff%2==1{
            return 0
        }       
        let neg=(diff/2) as usize; 
        for  i in 0..n+1  {           
            dp.push(Vec::new());
            dp[i].resize(neg+1, 0);
        }

        dp[0][0]=1;

        for i in 0..n  {
            dp.push(vec![]);
            let num =nums[i] as usize;
            for j in 0.. neg+1{
                dp[i+1][j]=dp[i][j];
                if j>=num{
                    dp[i+1][j] +=dp[i][j-num];
                }
            }   
        }

        dp[n][neg]
    }
}

#[test]
fn find_target_sum_ways_test(){
    assert_eq!( Solution::find_target_sum_ways(vec![1,1,1,1,1], 3),5)
}