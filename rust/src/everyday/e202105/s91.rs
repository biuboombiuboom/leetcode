struct Solution {}
impl Solution {
    pub fn num_decodings(s: String) -> i32 {
        let mut decodings: Vec<i32> = Vec::new();
        let mut num = s[0..1].parse::<i32>().unwrap();
        if num > 0 && num < 10 {
            decodings.push(1);
        } else {
            return 0;
        }
        for i in 1..s.len() {
            num = s[i..i + 1].parse::<i32>().unwrap();
           
            let double_num= s[i-1..i + 1].parse::<i32>().unwrap();
            if num == 0 {
                if double_num==0||double_num>27{
                    return 0
                }
                if i < 2 || decodings[i - 2] == decodings[i - 1] {
                    decodings.push(decodings[i - 1]);
                } else {
                    if i>1{
                        decodings.push(decodings[i - 2] );
                    }else{
                        decodings.push(decodings[i - 1] );
                    }
                }
                println!("{}-{}--{}",i,   num, decodings[decodings.len() - 1]);
                continue;
            }            
            if double_num>9&&double_num<27{
                if i>1{
                    decodings.push(decodings[i - 1] + decodings[i - 2]);
                }else{
                    decodings.push(decodings[i - 1] + 1);
                }
                
            }else{
                decodings.push(decodings[i - 1] );
            }     
            println!("{}-{}--{}",i,   num, decodings[decodings.len() - 1]);
        }
        decodings[decodings.len() - 1]
    }
}

pub fn run() {
    println!("{}", Solution::num_decodings(String::from("210")))
}
