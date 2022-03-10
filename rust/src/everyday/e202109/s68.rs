use std::iter::FromIterator;
struct Solution {}

impl Solution {
    pub fn full_justify(words: Vec<String>, max_width: i32) -> Vec<String> {
        let mut ans: Vec<String> = Vec::new();
        let max_width = max_width as usize;
        let mut part_words: Vec<String> = Vec::new();
        let mut part_len = 0;
        for word in words {
            let word_len = word.len();
            if part_len + word_len + part_words.len() <= max_width {
                part_len += word_len;
            } else if part_words.len() == 1 {
                let mut part_ans = part_words.join(" ");
                let l = part_ans.len();
                if l < max_width {
                    part_ans.push_str(
                        String::from_iter(std::iter::repeat(" ").take(max_width - l)).as_str(),
                    );
                }
                ans.push(part_ans);
                part_words.clear();
                part_len = word_len;
            } else {
                let space_word_len = max_width - part_len;
                let space_count = space_word_len / (part_words.len() - 1);
                let last = space_word_len - (part_words.len() - 1) * space_count;
                for i in 0..part_words.len() - 1 {
                    let mut c = space_count;
                    if i + 1 <= last {
                        c = c + 1;
                    }
                    part_words[i]
                        .push_str(String::from_iter(std::iter::repeat(" ").take(c)).as_str());
                }

                let part_ans = part_words.join("");
                ans.push(part_ans);
                part_words.clear();
                part_len = word_len;
            }
            part_words.push(word);
        }
        if part_words.len() > 0 {
            let mut word = part_words.join(" ");
            let l = word.len();
            if l < max_width {
                word.push_str(
                    String::from_iter(std::iter::repeat(" ").take(max_width - l)).as_str(),
                );
            }
            ans.push(word);
        }
        return ans;
    }
}

#[test]
fn full_justify_test() {
    assert_eq!(
        vec!["What   must   be", "acknowledgment  ", "shall be        "],
        Solution::full_justify(
            vec![
                "What".to_string(),
                "must".to_string(),
                "be".to_string(),
                "acknowledgment".to_string(),
                "shall".to_string(),
                "be".to_string()
            ],
            16,
        )
    )
}

// vec![
//     "This".to_string(),
//     "is".to_string(),
//     "an".to_string(),
//     "example".to_string(),
//     "of".to_string(),
//     "text".to_string(),
//     "justification.".to_string(),
// ],
