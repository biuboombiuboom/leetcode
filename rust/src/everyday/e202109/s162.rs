struct Solution;
impl Solution {
    pub fn find_peak_element(nums: Vec<i32>) -> i32 {
        let n = nums.len() as i32;
        if n == 1 {
            return 0;
        }
        let get = |i| {
            if i == -1 || i == n {
                return i32::MIN;
            } else {
                return nums[i as usize];
            }
        };
        let mut left: i32 = 0;
        let mut right: i32 = n - 1;
        loop {
            let mid = (left + right) / 2;
            if get(mid - 1) < get(mid) && get(mid) > get(mid + 1) {
                return mid;
            }
            if get(mid) < get(mid + 1) {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
}

#[test]
fn find_peak_element_test() {
    assert_eq!(
        Solution::find_peak_element(vec![-2147483648, 2147483646, 2147483647]),
        2
    )
}
