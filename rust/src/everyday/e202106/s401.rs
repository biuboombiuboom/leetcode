struct Solution;
impl Solution {
    pub fn read_binary_watch(turned_on: i32) -> Vec<String> {
        let mut ans: Vec<String> = Vec::new();
        let h: u32 = 12;
        let m: u32 = 60;
        let turned_on = turned_on as u32;
        for i in 0..h {
            for j in 0..m {
                if i.count_ones() + j.count_ones() == turned_on {
                    ans.push(format!("{}:{:<02}", i,j));
                }
            }
        }
        ans
    }
}
//h 8 4 2 1

//m 32 16 8 4 2 1

// 输入：turnedOn = 1
// 输出：["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
// 示例 2：

// 输入：turnedOn = 9
// 输出：[]

#[test]
fn read_binary_watch_test() {
    assert_eq!(Solution::read_binary_watch(9), vec!["".to_string()])
}
