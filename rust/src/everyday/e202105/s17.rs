pub fn letter_combinations(digits: String) -> Vec<String> {
    let mut ans: Vec<String> = Vec::new();
    let words = ["abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"];
    let len = digits.len();
    let mut i = 0;
    while i < len {
        let num: usize = digits[i..i + 1].parse().expect("");
     
        let word = &words[num - 2];
        let mut new_words: Vec<String> = Vec::new();
        for c in word.chars() {
            if ans.len() == 0 {
                new_words.push(String::from(c));
            } else {
                for ans_str in &mut ans {
                    let new_str = format!("{}{}", ans_str, c);
                    new_words.push(new_str);
                }
            }
        }
        i +=  1;
        
        ans = new_words;
    }
    ans
}
