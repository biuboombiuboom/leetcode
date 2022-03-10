use core::num;

use crate::utils::common::max;

struct Solution;
impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut max_jump = 0;
        for (i, &num) in nums.iter().enumerate() {
            if max_jump < i {
                return false;
            }
            max_jump = max_jump.max(i + num as usize)
        }

        return max_jump >= nums[nums.len() - 1] as usize;
    }
}

#[test]
fn can_jump_test() {
    let jump = vec![3, 2, 1, 0, 4];
    assert_eq!(Solution::can_jump(jump), false);
}
