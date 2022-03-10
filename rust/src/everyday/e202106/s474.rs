
struct Solution;
impl Solution {
    pub fn find_max_form(strs: Vec<&str>, m: i32, n: i32) -> i32 {
        let target = m + n;
        let m = m as usize;
        let n = n as usize;
        let l = strs.len();

        let mut dp: Vec<Vec<Vec<i32>>> = Vec::new();
        for i in 0..l + 1 {
            dp.push(Vec::new());
            for j in 0..m + 1 {
                dp[i].push(Vec::new());
                dp[i][j].resize(n + 1, 0);
            }
        }


        for i in 0..strs.len() {
           let zero_count= strs[i].matches("0").count();
           let one_count=strs[i].len()-zero_count;
            for j in 0..m+1{
                for k in 0..n+1{
                    dp[i+1][j][k] = dp[i][j][k];
                    if j >= zero_count && k >= one_count {
                        dp[i+1][j][k] = std::cmp::max(dp[i+1][j][k], dp[i][j-zero_count][k-one_count]+1)
                    }
                }
            }

        }
        return  dp[l][m][n];
    }
}

#[test]
fn find_max_form_test(){
    assert_eq!(Solution::find_max_form(vec!["10", "0001", "111001", "1", "0"], 5, 3),4);
}