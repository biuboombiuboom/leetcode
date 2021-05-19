use std::collections::HashMap;
struct Solution{}

impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut map:HashMap<i32,i32>=HashMap::new();
        for elem in nums {
            if let Some(count)=map.get(&elem){
                map.insert(elem,count+1);
            }else{
                map.insert(elem,1);
            }            
        }

        let mut single=0;
        for (k,v) in map{
            if v==1{
                single= k;
                break
            }
        }

        return single
    }
}

pub fn run(){
    println!("{}",Solution::single_number(vec! [30000,500,100,30000,100,30000,100]));
}