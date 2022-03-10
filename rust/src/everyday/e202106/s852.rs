struct Solution;

impl Solution {
    pub fn peak_index_in_mountain_array(arr: Vec<i32>) -> i32 {
        let mut top = arr[0];
        let mut top_index = 0;
        for i in 0..arr.len() {
            if arr[i] < top {
                break;
            }
            top = arr[i];
            top_index = i as i32;
        }
        return top_index;
    }
}

#[test]
fn peak_index_in_mountain_array_test() {
    assert_eq!(Solution::peak_index_in_mountain_array(vec![24, 69, 19]), 1)
}
