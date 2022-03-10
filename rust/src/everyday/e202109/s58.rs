struct Solution;
impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        let ss = s.trim_end().split(" ");
        return ss.last().unwrap().len() as i32;
    }
}
