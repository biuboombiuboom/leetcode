struct Solution;
impl Solution {
    pub fn stone_game(piles: Vec<i32>) -> bool {
        let n=piles.len();
        let mut dp=vec![vec![0;n];n];
        for i in 0..n{
            dp[i][i]=piles[i];
        }
        for i in 0..n{
            for j in i+1..n{
                if i==j{
                    dp[i][j]=piles[i];
                }
            }
        }
        false
    }
}

#[test]
fn stone_game_test() {
    assert_eq!(Solution::stone_game(vec![1, 2, 3, 4, 5]), false)
}
