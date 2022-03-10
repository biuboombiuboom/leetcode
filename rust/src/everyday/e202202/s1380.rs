use std::i32::{MAX, MIN};

impl Solution {
    pub fn lucky_numbers(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let mut lucky_numbers = Vec::new();
        let mut row_min = Vec::new();
        let mut col_max = Vec::new();
        let m = matrix.len();
        let n = matrix[0].len();
        row_min.resize(matrix.len(), MAX);
        col_max.resize(matrix[0].len(), MIN);
        for i in 0..m {
            for j in 0..n {
                let num = matrix[i][j];
                row_min[i] = row_min[i].min(num);
                col_max[j] = col_max[j].max(num);
            }
        }

        for i in 0..m {
            for j in 0..n {
                let num = matrix[i][j];
                if num == row_min[i] && num == col_max[j] {
                    lucky_numbers.push(num);
                }
            }
        }

        return lucky_numbers;
    }
}

struct Solution;

#[test]
fn lucky_numbers_test() {
    let input = vec![vec![1, 10, 4, 2], vec![9, 3, 8, 7], vec![15, 16, 17, 12]];
    let output = vec![12];
    assert_eq!(Solution::lucky_numbers(input), output);
}
