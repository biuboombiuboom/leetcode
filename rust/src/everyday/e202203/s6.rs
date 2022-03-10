struct Solution;
impl Solution {
    pub fn convert(s: String, num_rows: i32) -> String {
        let n = s.len();
        let num_rows = num_rows as usize;
        if n == 1 || num_rows == 1 {
            return s;
        }

        let mut rows = Vec::new();
        rows.resize_with(num_rows, || String::default());
        let mut cur = 0;
        let mut down = false;
        for i in 0..n {
            let s1 = &s[i..i + 1];
            rows[cur].push_str(s1);
            if cur == 0 || cur == num_rows - 1 {
                down = !down
            }
            if down {
                cur += 1
            } else {
                cur -= 1
            }
        }
        return rows.join("");
    }
}

#[test]
fn convert_test() {
    let input = "PAYPALISHIRING".to_string();
    let out = "PINALSIGYAHRPI";
    assert_eq!(Solution::convert(input, 4), out)
}
