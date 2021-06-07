impl Solution {
    pub fn xor_operation(n: i32, start: i32) -> i32 {
        let mut ans=start;
        for i in 1..n{
            let cur=start+2*i;
            ans=ans^cur;           
        }
        ans
    }
}

struct Solution{}

pub fn run(){
    println!("{}",Solution::xor_operation(5, 0));
}