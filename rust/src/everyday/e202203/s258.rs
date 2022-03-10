struct Solution;
impl Solution {
    pub fn add_digits(num: i32) -> i32 {
        let mut num = num;
        while num > 9 {
            let num_str = num.to_string();
            num = 0;
            for i in 0..num_str.len() {
                num += num_str[i..i + 1].parse::<i32>().unwrap();
            }
        }
        return num;
    }
}

#[test]
fn add_digits_test() {
    assert_eq!(Solution::add_digits(38), 2)
}
