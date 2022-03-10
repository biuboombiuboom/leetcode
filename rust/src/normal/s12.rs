fn int_to_roman(num: i32) -> String {
    let words = vec!["I", "V", "X", "L", "C", "D", "M"];
    let mut ans = String::new();
    let num_str = num.to_string();
    let mut i = num_str.len();
    while i > 0 {
        let mut n: usize = num_str[i - 1..i].parse().expect("");
        let bit = num_str.len() - i + 1;
        let mut bit_word = String::new();
        while n > 0 {
            if n == 9 {
                bit_word = String::from(words[bit * 2 - 2]) + &String::from(words[bit * 2]);
                n = 0;
            } else if n == 4 {
                bit_word = String::from(words[bit * 2 - 2]) + &String::from(words[bit * 2 - 1]);
                n = 0;
            } else if n >= 5 {
                bit_word = String::from(words[bit * 2 - 1]);
                n -= 5;
            } else {
                bit_word = bit_word + &String::from(words[(bit - 1) * 2]);
                n -= 1;
            }
        }
        ans = bit_word + &ans;
        i -= 1;
    }
    ans
}
