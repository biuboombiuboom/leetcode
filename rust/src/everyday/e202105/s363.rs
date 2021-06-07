use std::collections::BTreeSet;
struct Solution {}
impl Solution {
    pub fn max_sum_submatrix(matrix: Vec<Vec<i32>>, k: i32) -> i32 {
        let mut ans=i32::MIN;
        let m = matrix.len();
        let n = matrix[0].len();

        for i in 0..m {
            let mut sums = Vec::with_capacity(n);
            sums.resize_with(n, || 0);
            // sums.set_len(n);
            for j in i..m {
                for c in 0..n {
                    sums[c] += matrix[j][c];
                }
                let mut sum_set = BTreeSet::new();
                sum_set.insert(0);
                let mut s = 0;
                for sum in &sums {
                    s += sum;
                    if let Some(cc)=sum_set.range((s - k)..).next(){
                        ans=std::cmp::max(ans,s-cc);
                    }
                    sum_set.insert(s);
                }
            }
        }

        return ans;
    }
}

pub fn run() {
    println!("{}", Solution::max_sum_submatrix(vec![vec![1,0,1],vec![0,-2,3]], 2));
}
