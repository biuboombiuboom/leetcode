struct Solution;
impl Solution {
    pub fn sub_array_ranges(nums: Vec<i32>) -> i64 {
        let mut range_sum: i64 = 0;
        let n = nums.len();
        let mut sub_arrays = Vec::new();
        nums.iter().for_each(|&num| sub_arrays.push(vec![num, num]));
        for sub_len in 2..n + 1 {
            for i in 0..n - sub_len + 1 {
                let min = sub_arrays[i][0];
                let max = sub_arrays[i][1];
                let num = nums[i + sub_len - 1];
                sub_arrays[i][0] = min.min(num);
                sub_arrays[i][1] = max.max(num);
                range_sum += (sub_arrays[i][1] - sub_arrays[i][0]) as i64;
            }
        }

        return range_sum;
    }
}

#[test]
fn sub_array_ranges_test() {
    let input = vec![1, 5, 3, 8];
    // 1,5 4
    // 5,3 2
    // 3,8 5 11
    // 1,5,3 4 15
    // 5,3,8 5 21
    // 1,5,3,8 7
    let out = 27;
    assert_eq!(Solution::sub_array_ranges(input), out)
}
