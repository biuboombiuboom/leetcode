
struct Solution;
impl Solution {
    pub fn largest_number(cost: Vec<i32>, target: i32) -> String {
        let n = 9;
        let target = target as usize;
        let mut dp = Vec::new();
        let mut from = Vec::new();
        from.resize(10, vec![0; target + 1]);
        dp.resize(10, vec![i32::MIN; target + 1]);
        dp[0][0] = 0;
        for i in 0..n  {
            let c = cost[i ] as usize;
            for j in 0..target + 1 {
                if j < c {
                    dp[i + 1][j] = dp[i][j];
                    from[i + 1][j] = j;
                } else {
                    if dp[i][j] > dp[i+1][j-c]+1 {
                        dp[i+1][j] = dp[i][j];
                        from[i+1][j] = j
                    } else {
                        dp[i+1][j] = dp[i+1][j-c] + 1;
                        from[i+1][j] = j - c
                    }    
                }
            }
        }
        if dp[9][target] < 0 {
            return "0".to_string();
        }
        let mut i=9 as usize;
        let mut j=target;
        let mut ans=String::default();
        while i>0{
            if j==from[i][j]{
                i-=1;
            }else{
                ans.push_str(i.to_string().as_str());       
                j = from[i][j]       
            }
        }
        return ans;
    }
}

#[test]
fn largest_number_test() {
    assert_eq!(
        Solution::largest_number(vec![4, 3, 2, 5, 6, 7, 2, 5, 5], 9),
        "7772".to_string()
    )
}
