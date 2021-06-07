struct Solution {}

impl Solution {
    pub fn can_cross(stones: Vec<i32>) -> bool {
        let n = stones.len();
        let mut dp: Vec<Vec<bool>> = Vec::new();
        for i in 0..n {
            dp.push(Vec::with_capacity(n));
            dp[i].resize_with(n, || false);
        }
        dp[0][0] = true;

        for i in 1..n {
            if stones[i] - stones[i - 1] > i as i32 {
                return false;
            }
        }

        for i in 1..n {
            let mut jj = (i - 1) as i32;
            while jj >= 0 {
                let j=jj as usize;
                let k = (stones[i] - stones[j] )as usize;
                if k > j + 1 {
                    break;
                }
                dp[i][k] = dp[j][k - 1] || dp[j][k] || dp[j][k + 1];
                if i == n - 1 && dp[i][k] {
                    return true;
                }               
                jj-=1;
            }
        }
        return false;
    }
}

pub fn run() {
    println!("{}", Solution::can_cross(vec![0, 1, 3, 5, 6, 8, 12, 17]))
}
