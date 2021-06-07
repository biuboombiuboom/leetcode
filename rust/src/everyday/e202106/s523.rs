use std::{collections::HashMap,};

struct Solution;
impl Solution {
    pub fn check_subarray_sum(nums: Vec<i32>, k: i32) -> bool {       
        let n = nums.len() as i32;
        if n<2{
            return false
        }
        let mut rem_map=HashMap::new();    
        rem_map.insert(0, -1);

        let mut remainder = 0;
        for i in 0..n { 
            let num=nums[i as usize];
            remainder = (remainder + num) % k;
            if let Some(prev_index)=rem_map.get(&remainder){
                if i-prev_index >= 2 {
                    return true
                }
            }else{
                rem_map.insert(remainder, i);
            }
        }

        return false;
    }
   
}


#[test]
fn check_subarray_sum_test(){
    assert_eq!(Solution::check_subarray_sum(vec![23,2,4,6,6],7),true)
}