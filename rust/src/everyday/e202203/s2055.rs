struct Solution;
impl Solution {
    pub fn plates_between_candles(s: String, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ans = Vec::new();
        let mut sum_p = Vec::new();
        let n = s.len();
        let mut l = n;
        let mut r = n;
        let mut left = Vec::new();
        let mut right = Vec::new();
        let mut sum = 0;
        right.resize_with(n, || n);
        for (i, &b) in s.as_bytes().iter().enumerate() {
            if b == b'*' {
                sum += 1;
            } else {
                l = i;
            }
            if s.as_bytes()[n - 1 - i] == b'|' {
                r = n - 1 - i
            }
            sum_p.push(sum);
            left.push(l);
            right[n - 1 - i] = r;
        }

        for i in 0..queries.len() {
            let query = &queries[i];
            let x = right[query[0] as usize];
            let y = left[query[1] as usize];

            if x < n && y < n && x < y {
                ans.push(sum_p[y as usize] - sum_p[x as usize])
            } else {
                ans.push(0);
            }
        }

        return ans;
    }
}

#[test]
fn plates_between_candles_test() {
    let s = String::from("***|**|*****|**||**|*");
    let queries = vec![
        vec![1, 17],
        vec![4, 5],
        vec![14, 17],
        vec![5, 11],
        vec![15, 16],
    ];
    let ans = vec![9, 0, 0, 0, 0];

    assert_eq!(Solution::plates_between_candles(s, queries), ans);
}
