use std::cmp::Ordering;

struct Solution;
impl Solution {
    pub fn find_longest_word(s: String, dictionary: Vec<String>) -> String {
        let dn = dictionary.len();
        let mut ans = "".to_string();
        for i in 0..dn {
            let word = dictionary[i].clone();
            if word.len() > s.len() {
                continue;
            }
            let mut is_match = true;
            if let Some(index) = s.find(word.as_bytes()[0] as char) {
                let mut pre_index = index;
                for c in word[1..].chars() {
                    if s.len() - 1 == pre_index {
                        is_match = false;
                        break;
                    }
                    if let Some(i) = s[pre_index + 1..].find(c) {
                        pre_index = pre_index + 1 + i
                    } else {
                        is_match = false;
                        break;
                    }
                }
            } else {
                is_match = false
            }

            if is_match {
                let wn = word.len();
                if ans.len() < wn {
                    ans = word.clone()
                } else if ans.len() == wn {
                    if Ordering::Greater == ans.cmp(&word) {
                        ans = word
                    }
                }
            }
        }
        return ans;
    }
}

#[test]
fn find_longest_word_test() {
    let s = String::from("abpcplea");

    let dictionary = vec!["ale", "apple", "monkey", "plea"];
    let dictionary = dictionary.iter().map(|s| s.to_string()).collect();

    let ans = String::from("apple");
    assert_eq!(Solution::find_longest_word(s, dictionary), ans);
}
