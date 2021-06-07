struct Solution;
impl Solution {
    pub fn count_triplets(arr: Vec<i32>) -> i32 {
        let mut xors:Vec<i32> =Vec::new();
        xors.push(arr[0]);
        let n =arr.len();
        for i in 0..n{
            xors.push(xors[i]^arr[i])
        }
        println!("{:?}",xors);
        let mut ans=0;
        for i in 0..n{
            for j in i+1..n{       
                for k in j..n{
                    if xors[i]==xors[k+1]{
                        ans+=1;
                    }
                }
            }                  
        }
        ans
    }
}

pub fn run(){
    println!("{}",Solution::count_triplets(vec![2,3,1,6,7]))
}