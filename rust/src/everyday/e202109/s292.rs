struct Solution {}
impl Solution {
    pub fn can_win_nim(n: i32) -> bool {
        return n % 4 != 0;
    }
}

#[test]
fn can_win_nim_test() {
    let n = 0;
    let ans = Solution::can_win_nim(n);
    let expect = true;
    assert_eq!(ans, expect)
}
