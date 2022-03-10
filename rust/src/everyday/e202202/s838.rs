use std::iter::FromIterator;

impl Solution {
    pub fn push_dominoes(dominoes: String) -> String {
        let mut ans = String::default();
        let mut pre_r: Option<usize> = None;
        for (i, &b) in dominoes.as_bytes().iter().enumerate() {
            if b == b'R' {
                if let Some(_) = pre_r {
                    ans.push('R')
                }
                if i > ans.len() {
                    let stand_iters = std::iter::repeat(".").take(i - ans.len());
                    let stand_string = String::from_iter(stand_iters);
                    ans.push_str(stand_string.as_str());
                }
                pre_r = Some(i);

                continue;
            }
            if b == b'L' {
                if let Some(r_index) = pre_r {
                    let range = i - r_index + 1;
                    let r_iters = std::iter::repeat("R").take(range / 2);
                    let r_string = String::from_iter(r_iters);
                    ans.push_str(r_string.as_str());
                    if range % 2 != 0 {
                        ans.push('.')
                    }
                    let l_iters = std::iter::repeat('L').take(range / 2);
                    let l_string = String::from_iter(l_iters);
                    ans.push_str(l_string.as_str());
                    pre_r = None
                } else {
                    let range = i - ans.len() + 1;
                    let l_iters = std::iter::repeat('L').take(range);
                    let l_string = String::from_iter(l_iters);
                    ans.push_str(l_string.as_str());
                }
            }
        }

        if dominoes.len() > ans.len() {
            if let Some(_) = pre_r {
                let range = dominoes.len() - ans.len();
                let r_iters = std::iter::repeat("R").take(range);
                let r_string = String::from_iter(r_iters);
                ans.push_str(r_string.as_str());
            } else {
                ans.push_str(&dominoes[ans.len()..]);
            }
        }
        return ans;
    }
}

struct Solution;

#[test]
fn push_dominoes() {
    //.L.R...LR..L..
    let input = "R.R.L".to_string();
    let output = "RRR.L";
    assert_eq!(Solution::push_dominoes(input), output)
}
