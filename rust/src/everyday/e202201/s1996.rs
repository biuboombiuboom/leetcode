use std::cmp::Ordering;

impl Solution {
    pub fn number_of_weak_characters(properties: Vec<Vec<i32>>) -> i32 {
        let mut ans = 0;
        let n = properties.len();
        let mut skip_map = Vec::new();
        skip_map.resize(n, false);
        let mut properties = properties;

        for i in 0..n {
            if skip_map[i] {
                continue;
            }
            let mut is_weak = false;
            for j in i + 1..n {
                if skip_map[j] {
                    continue;
                }
                if !is_weak
                    && properties[i][0] < properties[j][0]
                    && properties[i][1] < properties[j][1]
                {
                    ans += 1;
                    is_weak = true;
                }
                if properties[i][0] > properties[j][0] && properties[i][1] > properties[j][1] {
                    ans += 1;
                    skip_map[j] = true;
                }
            }
        }
        return ans;
    }
}

struct Solution;

#[test]
fn number_of_weak_characters_test() {
    let input = vec![
        vec![7, 7],
        vec![1, 2],
        vec![9, 7],
        vec![7, 3],
        vec![3, 10],
        vec![9, 8],
        vec![8, 10],
        vec![4, 3],
        vec![1, 5],
        vec![1, 5],
    ];
    assert_eq!(Solution::number_of_weak_characters(input), 6);
}
