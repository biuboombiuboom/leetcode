impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        if n == 1 {
            return nums[0];
        }
        let mut dp = Vec::new();
        dp.push(nums[0]);
        dp.push(max(nums[0], nums[1]));
        for i in 2..n {
            dp.push(max(dp[i - 2] + nums[i], dp[i - 1]));
        }
        return dp[n - 1];
    }
}

fn max(a: i32, b: i32) -> i32 {
    if a > b {
        return a;
    }
    return b;
}

struct Solution;

#[test]
fn rob_test() {
    let input = vec![2, 7, 9, 3, 1];
    assert_eq!(Solution::rob(input), 12)
}
