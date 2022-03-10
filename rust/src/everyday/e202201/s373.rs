use std::{collections::HashMap, num};

struct Solution;
impl Solution {
    pub fn k_smallest_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        let mut pool: HashMap<i32, Vec<i32>> = HashMap::new();
        let mut ans = Vec::new();
        let m = nums1.len();
        let n = nums2.len();
        for i in 0..m {
            let x = nums1[i];
            for j in i + 1..m {
                let y = nums1[j];
            }
            for k in 0..n {}
        }
        for i in 0..n {
            for j in i + 1..n {}
        }

        return ans;
    }
}
