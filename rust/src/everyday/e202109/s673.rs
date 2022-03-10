use core::num;

use crate::utils::common::max;

struct Solution;
impl Solution {
    pub fn find_number_of_lis(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut dp = vec![0; n];
        let mut cnt = vec![0; n];
        let mut max_len = 0;
        let mut ans = 0;
        for i in 0..nums.len() {
            dp[i] = 1;
            cnt[i] = 1;
            let x = nums[i];
            for j in 0..i {
                let y = nums[j];
                if x > y {
                    if dp[j] + 1 > dp[i] {
                        dp[i] = dp[j] + 1;
                        cnt[i] = cnt[j];
                    } else if dp[j] + 1 == dp[i] {
                        cnt[i] += cnt[j];
                    }
                }
            }
            if dp[i] > max_len {
                max_len = dp[i];
                ans = cnt[i];
            } else if dp[i] == max_len {
                ans += cnt[i];
            }
        }
        return ans;
    }
}

#[test]
fn find_number_of_lis_test() {
    let left = Solution::find_number_of_lis(vec![2, 2, 2, 2, 2]);
    let right = 5;
    assert_eq!(left, right);
}
