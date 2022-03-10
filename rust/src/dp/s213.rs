impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        if n == 1 {
            return nums[0];
        }
        let mut dp = Vec::new();
        dp.push(vec![0, nums[0]]);
        dp.push(vec![nums[1], nums[0]]);
        for i in 2..n - 1 {
            let a = max(dp[i - 2][0] + nums[i], dp[i - 1][0]);
            let b = max(dp[i - 2][1] + nums[i], dp[i - 1][1]);
            dp.push(vec![a, b]);
        }
        if n > 2 {
          dp.push(vec![
            max(dp[n - 3][0] + nums[n - 1], dp[n - 2][0]),
            max(dp[n - 3][1], dp[n - 2][1]),
        ]);

        }
    
        return max(dp[n - 1][0], dp[n - 1][1]);
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
    let input = vec![0, 0];
    assert_eq!(Solution::rob(input), 0)
}
