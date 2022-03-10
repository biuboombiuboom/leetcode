struct Solution {}
impl Solution {
    pub fn number_of_boomerangs(points: Vec<Vec<i32>>) -> i32 {
        use std::collections::HashMap;
        let l = points.len();
        let mut ans = 0;
        for i in 0..l {
            let mut disc: HashMap<i32, i32> = HashMap::new();
            let pi = &points[i];
            for j in 0 ..l {
                let pj = &points[j];
                let dis = (pi[0]  -pj[0]) * (pi[0]  -pj[0]) + (pi[1] - pj[1]) * (pi[1] - pj[1]);
                let c = disc.entry(dis).or_insert(0);
                *c += 1;
            }
            for (&d, &c) in &disc {
                if d > 0 {
                    ans += c * (c - 1)
                }
            }
        }
        return ans;
    }
}

#[test]
fn number_of_boomerangs_test() {
    assert_eq!(
        Solution::number_of_boomerangs(vec![vec![0, 0], vec![1, 0], vec![2, 0]]),
        2
    )
}
