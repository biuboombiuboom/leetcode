struct Solution;

impl Solution {
    pub fn get_sum(a: i32, b: i32) -> i32 {
        let mut a = a;
        let mut b = b;
        while b != 0 {
            let carry = (a & b) << 1;
            a = a ^ b;
            b = carry;
        }
        return a;
    }
}

#[test]
fn get_sum_test() {
    assert_eq!(Solution::get_sum(2, 3), 5);
}
