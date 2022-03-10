struct Solution {}

// impl Solution {
//     pub fn judge_square_sum(c: i32) -> bool {
//         let mut i = 0;
//         while i * i <= c {
//             let f = (c - i * i) as f64;
//             let rt = f.sqrt();
//             if rt == rt.floor() {
//                 return true;
//             }
//             i += 1;
//         }
//         return false;
//     }
// }

// func judgeSquareSum(c int) bool {
//     left, right := 0, int(math.Sqrt(float64(c)))
//     for left <= right {
//         sum := left*left + right*right
//         if sum == c {
//             return true
//         } else if sum > c {
//             right--
//         } else {
//             left++
//         }
//     }
//     return false
// }

impl Solution {
    pub fn judge_square_sum(c: i32) -> bool {
        let mut left = 0;
        let f = c as f64;
        let mut right = f.sqrt() as i32;
        while left <= right {
            let sum = left * left + right * right;
            if sum == c {
                return true;
            } else if sum > c {
                right -= 1;
            } else {
                left += 1;
            }
        }

        return false;
    }
}

pub fn run() {
    println!("{}", Solution::judge_square_sum(999999999));
}
