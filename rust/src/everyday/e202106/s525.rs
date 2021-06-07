struct Solution;
impl Solution {
    pub fn find_max_length(nums: Vec<i32>) -> i32 {
        let n = nums.len();   
        let mut map = std::collections::HashMap::new();
        map.insert(0, -1);
        let mut max_len = 0;
        let mut counter = 0;
        for i in 0..n {
            let num = nums[i];
            if num == 1 {
                counter += 1;
            } else {
                counter -= 1;
            }

            if let Some(prev_index) = map.get(&counter) {
                max_len = std::cmp::max(max_len, i as i32-prev_index)
            }else{
                map.insert(counter,i as i32);
            }
        }


        max_len
    }
}

#[test]
fn find_max_length_test() {
    assert_eq!(Solution::find_max_length(vec![1, 0, 0, 0, 0, 1, 1]), 4);
}
