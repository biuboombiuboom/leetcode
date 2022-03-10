struct Solution;

impl Solution {
    pub fn can_eat(candies_count: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<bool> {
        let n=queries.len();
        let mut ans=Vec::with_capacity(n);
        ans.resize(n, false);
        ans
    }
}