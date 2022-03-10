use core::num;

use crate::utils::common::max;

impl Solution {
    pub fn minimum_difference(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        nums.sort();
        let k = k as usize;
        let mut minmum = nums[k - 1] - nums[0];
        let mut i = 0;
        let mut j = k - 1;
        while j < nums.len() {
            if minmum > nums[j] - nums[i] {
                minmum = nums[j] - nums[i]
            }
            i += 1;
            j += 1;
        }
        return minmum;
    }
}

struct Solution;

#[test]
fn minimum_difference_test() {
    let input = vec![9, 4, 1, 7];
    assert_eq!(Solution::minimum_difference(input, 2), 2);
}
