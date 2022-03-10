struct Solution;
impl Solution {
    pub fn modify_string(s: String) -> String {
        let mut ans = String::default();
        let mark_ch = '?';
        let mut pre_ch = '?';
        for (i, &item) in s.as_bytes().iter().enumerate() {
            if item == b'?' {
                for b in 97..124 {
                    if ans.as_bytes()[i - 1] != b && s.as_bytes()[i + 1] != b {
                        ans.push(b as char);
                        break;
                    }
                }
            } else {
                ans.push(item as char);
            }
        }

        return ans;
    }
}

#[test]
fn modify_string_test() {
    let input = String::from("a?");
    let expect = String::from("ab");
    let output = Solution::modify_string(input);
    assert_eq!(expect, output)
}
