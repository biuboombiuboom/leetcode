struct Solution;

impl Solution {
    pub fn int_to_roman(num: i32) -> String {
        String::from("IVH")
    }
}

pub fn run(){
    println!("{}",Solution::int_to_roman(123)) ;
}