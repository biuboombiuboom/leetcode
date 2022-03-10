struct Solution;
impl Solution {
    pub fn nearest_palindromic(n: String) -> String {
        let num = n.parse::<i64>().unwrap();
        if n.len() == 1 || (n[0..1] == "1".to_string() && n[1..].parse::<i32>().unwrap() == 0) {
            return (num - 1).to_string();
        }
        if num == 11 {
            return "9".to_string();
        }
        let l = n.len() + 1;
        let left_num = n[0..l / 2].parse::<i64>().unwrap();
        let mut num1 = left_num.to_string();
        let mut num2 = (left_num + 1).to_string();
        let mut num3 = (left_num - 1).to_string();
        num1.push_str(
            num1[0..n.len() - l / 2]
                .to_string()
                .chars()
                .rev()
                .collect::<String>()
                .as_str(),
        );
        num2.push_str(
            num2[0..n.len() - l / 2]
                .to_string()
                .chars()
                .rev()
                .collect::<String>()
                .as_str(),
        );
        num3.push_str(
            num3[0..n.len() - l / 2]
                .to_string()
                .chars()
                .rev()
                .collect::<String>()
                .as_str(),
        );
        let num1 = num1.parse::<i64>().unwrap();
        let num2 = num2.parse::<i64>().unwrap();
        let num3 = num3.parse::<i64>().unwrap();

        let mut ans_num = 0;
        if (num3 - num).abs() < (ans_num - num).abs() && num3 - num != 0 {
            ans_num = num3;
        }
        if (num1 - num).abs() < (ans_num - num).abs() && num1 - num != 0 {
            ans_num = num1;
        }

        if (num2 - num).abs() < (ans_num - num).abs() && num2 - num != 0 {
            ans_num = num2;
        }

        return ans_num.to_string();
    }
}

#[test]
fn nearest_palindromic_test() {
    let input = "172938225610270464".to_string();
    let output = "77";

    assert_eq!(Solution::nearest_palindromic(input), output)
}
