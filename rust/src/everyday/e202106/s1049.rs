use std::usize;

struct Solution;
impl Solution {
    pub fn last_stone_weight_ii(stones: Vec<i32>) -> i32 {

        
        let mut sum = 0;
        let n = stones.len();
        let mut dp=Vec::new();
        for i in 0..n {
            sum += stones[i];
        }
        let m=sum as usize/2;
        for i in 0..n+1{
            dp.push(Vec::new());
            dp[i].resize(m+1, false)
        }
        dp[0][0]=true;
        for i in 0..n{
            let wight=stones[i] as usize;
            for j in 0..m+1{
                if j<wight{
                    dp[i+1][j]=dp[i][j]
                }else{
                    dp[i+1][j]=dp[i][j]||dp[i][j-wight]
                }
            }
        }

        for i in (0..m+1).rev() {
            if dp[n][i]{
                return sum-2*i as i32
            }
        }

        return 0


    }
}

#[test]
fn last_stone_weight_ii_test() {
    assert_eq!(Solution::last_stone_weight_ii(vec![2,7,4,1,8,1]), 1);
}
