use core::num;

impl Solution {
    pub fn single_non_duplicate(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut m = (n - 1) / 2;
        let mut l = 0;
        let mut r = n - 1;
        while r - l > 3 {
            let mid = nums[m];
            let left = nums[m - 1];
            let right = nums[m + 1];
            if mid > left && mid < right {
                return mid;
            }
            if mid == left {
                //目标数在右区
                if (m - l + 1) % 2 == 0 {
                    l = m + 1;
                } else {
                    r = m - 2;
                }
            }
            if mid == right {
                //目标数在右区
                if (r - m + 1) % 2 == 1 {
                    l = m + 2;
                } else {
                    r = m - 1;
                }
            }

            m = l + (r - l) / 2;
        }
        if nums[l] == nums[m] {
            return nums[r];
        }
        return nums[l];
    }
}

struct Solution;

#[test]
fn single_non_duplicate_test() {
    let input = vec![1, 1, 3, 3, 7, 7, 10, 11, 11];
    // 9/2=4 0 4 8
    //

    assert_eq!(Solution::single_non_duplicate(input), 10);
}
