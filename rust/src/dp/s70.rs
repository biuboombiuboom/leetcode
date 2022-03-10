impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        let n = n as usize;
        let mut ans = vec![];
        ans.push(1);
        ans.push(2);
        for i in 2..n {
            ans.push(ans[i - 2] + ans[i - 1]);
        }
        return ans[n - 1];
    }
}

struct Solution;

#[test]
fn climb_stairs_test() {
    assert_eq!(Solution::climb_stairs(5), 8);
}
