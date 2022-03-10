struct Solution {}
impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut rows = vec![vec![0; 9]; 9];
        let mut cols = vec![vec![0; 9]; 9];
        let mut subboxs = vec![vec![vec![0; 9]; 3]; 3];
        for i in 0..9 {
            for j in 0..9 {
                let ch = board[i][j];
                if ch == '.' {
                    continue;
                }
                let index = (ch.to_digit(10).unwrap() as usize) - 1;
                rows[i][index] += 1;
                cols[j][index] += 1;
                subboxs[i / 3][j / 3][index] += 1;
                if rows[i][index] != 1 || cols[j][index] != 1 || subboxs[i / 3][j / 3][index] != 1 {
                    return false;
                }
            }
        }
        return true;
    }
}

#[test]
fn is_valid_sudoku_test() {
    let board = [
        ["5", "3", ".", ".", "7", ".", ".", ".", "."],
        ["6", ".", ".", "1", "9", "5", ".", ".", "."],
        [".", "9", "8", ".", ".", ".", ".", "6", "."],
        ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
        ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
        ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
        [".", "6", ".", ".", ".", ".", "2", "8", "."],
        [".", ".", ".", "4", "1", "9", ".", ".", "5"],
        [".", ".", ".", ".", "8", ".", ".", "7", "9"],
    ]
    .iter()
    .map(|item| item.iter().map(|s| s.as_bytes()[0] as char).collect())
    .collect();

    let ans = Solution::is_valid_sudoku(board);
    assert_eq!(true, ans);
}
