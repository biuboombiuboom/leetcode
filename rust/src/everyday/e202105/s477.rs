use std::fmt::Binary;

struct Solution;
impl Solution {
    pub fn total_hamming_distance(nums: Vec<i32>) -> i32 {
        let n = nums.len() as i32;
        let mut ans = 0;

        for i in 0..30 {
            let mut c = 0;
            for num in nums.iter() {
                c += (num >> i) & 1;
            }

            ans += c * (n - c)
        }
        return ans;
    }
}

#[test]
fn total_hamming_distance_test() {
    assert_eq!(Solution::total_hamming_distance(vec![14, 4, 2]), 6);
}
