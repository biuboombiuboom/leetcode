impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let n = cost.len();
        let mut dp = Vec::new();
        dp.push(0); //0
        dp.push(0); //1
        for i in 2..n + 1 {
            dp.push(min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]))
        }

        return dp[n];
    }
}

fn min(a: i32, b: i32) -> i32 {
    if a < b {
        return a;
    }
    return b;
}

struct Solution;

#[test]
fn min_cost_climbing_stairs_test() {
    let input = vec![1, 100, 1, 1, 1, 100, 1, 1, 100, 1];
    assert_eq!(Solution::min_cost_climbing_stairs(input), 6);
}
