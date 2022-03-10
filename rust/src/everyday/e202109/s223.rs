struct Solution;
impl Solution {
    pub fn compute_area(
        ax1: i32,
        ay1: i32,
        ax2: i32,
        ay2: i32,
        bx1: i32,
        by1: i32,
        bx2: i32,
        by2: i32,
    ) -> i32 {
        use std::cmp::max;
        use std::cmp::min;
        let x1 = max(ax1, bx1);
        let y1 = max(ay1, by1);
        let x2 = min(ax2, bx2);
        let y2 = min(ay2, by2);
        let s = (x1 - x2) * (y1 - y2);
        let s1 = (ax1 - ax2) * (ay1 - ay2);
        let s2 = (bx1 - bx2) * (by1 - by2);
        if x2 - x1 > 0 && y2 - y1 > 0 {
            return s1 + s2 - s;
        }
        return s1 + s2;
    }
}

#[test]
fn compute_area_test() {
    assert_eq!(Solution::compute_area(-2, -2, 2, 2, 3, 3, 4, 4), 16);
}
