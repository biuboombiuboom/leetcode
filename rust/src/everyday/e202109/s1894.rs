struct Solution {}
impl Solution {
    pub fn chalk_replacer(chalk: Vec<i32>, k: i32) -> i32 {
        let mut k = k as i64;
        let mut sum_chalk: i64 = 0;
        for i in 0..chalk.len() {
            sum_chalk += chalk[i] as i64;
        }
        if k % sum_chalk == 0 {
            return 0;
        }
        k = k - (k / sum_chalk * sum_chalk);
        for i in 0..chalk.len() {
            if k < chalk[i] as i64 {
                return i as i32;
            }
            k = k - chalk[i] as i64;
        }
        k = k / sum_chalk;
        return 0;
    }
}

#[test]
fn chalk_replacer_test() {
    use std::io::Read;

    let mut file = std::fs::File::open(
        "E:\\workspace\\rust\\leetcode\\rust\\src\\everyday\\e202109\\input.txt",
    )
    .unwrap();
    let mut input_strfile = String::default();
    file.read_to_string(&mut input_strfile);
    let input = input_strfile
        .split(",")
        .map(|s| s.parse::<i32>().unwrap())
        .collect();
    assert_eq!(Solution::chalk_replacer(input, 539095482), 10737);
}
