struct  Solution;
impl Solution {
    pub fn num_squares(n: i32) -> i32 {
        let mut dp=Vec::new();
        let n=n as usize;
        dp.resize(n+1, 0);

        for i in 1..n+1{
            let mut j:usize=1;
            let mut min=i32::MAX;
            while j*j<=i{
                min=std::cmp::min(min,dp[i-j*j]);

                j+=1;
            }
            println!("{}",min);
            dp[i]=min+1;
        }               
        return dp[n]
    }
}

#[test]
fn num_squares_test(){
    assert_eq!(Solution::num_squares(12),3);
}