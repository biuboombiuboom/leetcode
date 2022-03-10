struct Solution;
impl Solution {
    pub fn max_depth(s: String) -> i32 {
        let mut depth = 0;
        let mut ans = 0;
        for b in s.bytes() {
            if b == b'(' {
                depth += 1
            }
            if b == b')' {
                depth -= 1
            }
            if depth > ans {
                ans = depth
            }
        }
        return ans;
    }
}

#[test]
fn max_depth_test() {
    let input = String::from("AB");
    assert_eq!(Solution::max_depth(input), 0)
}
