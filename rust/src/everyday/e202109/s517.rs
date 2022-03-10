use std::cmp::max;

struct Solution;
impl Solution {
    pub fn find_min_moves(machines: Vec<i32>) -> i32 {
        fn abs(sum0: i32) -> i32 {
            if sum0 < 0 {
                return sum0 * -1;
            }
            return sum0;
        }
        let n = machines.len();
        let mut sum = 0;
        for i in 0..n {
            sum += machines[i];
        }

        if sum % n as i32 != 0 {
            return -1;
        }
        let avg = sum / n as i32;
        let mut sum0 = 0;
        let mut ans = 0;
        for mut num in machines {
            num -= avg;
            sum0 += num;
            ans = max(ans, max(abs(sum0), num))
        }
        return ans;
    }
}

#[test]
fn find_min_moves_test() {
    assert_eq!(Solution::find_min_moves(vec![5, 2, 2]), 2);
}
