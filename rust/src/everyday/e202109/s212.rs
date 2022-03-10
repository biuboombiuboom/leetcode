struct Solution {}
impl Solution {
    pub fn find_words(board: Vec<Vec<char>>, words: Vec<String>) -> Vec<String> {
        let mut ans = vec![];
        let m = board.len();
        let n = board[0].len();
        struct Walk<'s> {
            w: &'s Fn(
                &Walk,
                &Vec<Vec<char>>,
                &mut Vec<Vec<bool>>,
                &String,
                usize,
                usize,
                usize,
            ) -> bool,
        }
        let walk = Walk {
            w: &|walk, board, marked, word, i, j, start| -> bool {
                let wc = word.as_bytes()[start] as char;
                let c = board[i][j];
                println!("{},{},{},{},{}", c, wc, i, j, start);
                if start == word.len() - 1 && c == wc {
                    return true;
                }
                if c == wc {
                    marked[i][j] = true;
                    //上
                    if i > 0
                        && !marked[i - 1][j]
                        && (walk.w)(walk, board, marked, word, i - 1, j, start + 1)
                    {
                        return true;
                    }
                    //下
                    if i < m - 1
                        && !marked[i + 1][j]
                        && (walk.w)(walk, board, marked, word, i + 1, j, start + 1)
                    {
                        return true;
                    }
                    //左
                    if j > 0
                        && !marked[i][j - 1]
                        && (walk.w)(walk, board, marked, word, i, j - 1, start + 1)
                    {
                        return true;
                    }
                    //右
                    if j < n - 1
                        && !marked[i][j + 1]
                        && (walk.w)(walk, board, marked, word, i, j + 1, start + 1)
                    {
                        return true;
                    }
                    marked[i][j] = false;
                }

                return false;
            },
        };
        for word in &words {
            let board_tmp = board.clone();
            let mut marked = vec![vec![false; n]; m];
            let mut find = false;
            for i in 0..m {
                for j in 0..n {
                    if (walk.w)(&walk, &board_tmp.clone(), &mut marked, word, i, j, 0) {
                        ans.push((*word).clone());
                        find = true;
                        break;
                    }
                }
                if find {
                    break;
                }
            }
        }

        return ans;
    }
}

#[test]
fn find_words_test() {
    let board = vec![
        ["o", "a", "b", "n"],
        ["o", "t", "a", "e"],
        ["a", "h", "k", "r"],
        ["a", "f", "l", "v"],
    ]
    .iter()
    .map(|v| v.iter().map(|s| s.as_bytes()[0] as char).collect())
    .collect();
    let words = vec!["oa", "oaa"].iter().map(|s| s.to_string()).collect();
    let ans: Vec<String> = vec!["oa", "oaa"].iter().map(|s| s.to_string()).collect();
    assert_eq!(Solution::find_words(board, words), ans)
}
