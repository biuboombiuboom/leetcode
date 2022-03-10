struct Solution;
impl Solution {
    pub fn find_ball(grid: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ans = Vec::new();
        let r_len = grid.len();
        let c_len = grid[0].len();
        for x in 0..c_len {
            let mut xi = x;
            let mut lock = false;
            for y in 0..r_len {
                let de = grid[y][xi];
                let new_xi = xi as i32 + de;
                if new_xi < 0 || new_xi >= c_len as i32 || de + grid[y][new_xi as usize] == 0 {
                    ans.push(-1);
                    lock = true;
                    break;
                }
                xi = new_xi as usize;
            }
            if !lock {
                ans.push(xi as i32)
            }
        }
        return ans;
    }
}

#[test]
fn find_ball_test() {
    let grid = vec![
        vec![1, 1, 1, -1, -1],
        vec![1, 1, 1, -1, -1],
        vec![-1, -1, -1, 1, 1],
        vec![1, 1, 1, 1, -1],
        vec![-1, -1, -1, -1, -1],
    ];
    let output = vec![1, -1, -1, -1, -1];
    assert_eq!(Solution::find_ball(grid), output)
}
