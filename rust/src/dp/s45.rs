struct Solution;
impl Solution {
    pub fn jump(nums: Vec<i32>) -> i32 {
        let mut dp = Vec::new();
        let n = nums.len();

        dp.push(0);
        for i in (0..n - 1).rev() {
            let num = nums[i] as usize;
            if num == 0 {
                dp.push(n + 1);
                continue;
            }
            if nums[i] as usize + i >= n {
                dp.push(1);
            } else {
                let mut sub_dp = dp[dp.len() - num..].to_vec();
                sub_dp.sort();
                if sub_dp[0] != n + 1 {
                    dp.push(sub_dp[0] + 1);
                } else {
                    dp.push(n + 1);
                }
            }
        }

        return dp[n - 1] as i32;
    }
}

#[test]
fn jump_test() {
    let input = vec![5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0];
    assert_eq!(Solution::jump(input), 3);
}
