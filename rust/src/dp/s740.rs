impl Solution {
    pub fn delete_and_earn(nums: Vec<i32>) -> i32 {
        let mut max_val = 0;
        for num in nums.clone() {
            max_val = max_val.max(num);
        }
        let mut sum = Vec::new();
        sum.resize(max_val as usize + 1, 0);
        for num in nums {
            sum[num as usize] += num;
        }
        return rob(sum);
    }
}

fn rob(nums: Vec<i32>) -> i32 {
    let mut first = nums[0];
    let mut second = nums[0].max(nums[1]);
    for i in 2..nums.len() {
        let temp = first + nums[i];
        first = second;
        second = second.max(temp);
    }
    return second;
}

struct Solution;

#[test]
fn delete_and_earn_test() {
    let input = vec![1, 6, 3, 3, 8, 4, 8, 10, 1, 3];
    assert_eq!(Solution::delete_and_earn(input), 43);
}
