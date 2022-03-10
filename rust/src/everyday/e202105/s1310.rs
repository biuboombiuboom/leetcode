impl Solution {
    pub fn xor_queries(arr: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ans:Vec<i32>=Vec::new();
        let n=queries.len();
        for i in 0..n {
            let query=&queries[i];
            let start=query[0];
            let end=query[1];
            let mut xor=arr[start as usize];
            for j in start+1 ..end+1{
                let cur=arr[j as usize];
                xor=xor^cur;
            }
            ans.push(xor);
        }
        ans
    }
}

struct Solution{}

pub fn run(){
    println!("{:?}", Solution::xor_queries(vec![1,3,4,8], vec![vec![0,1],vec![1,2],vec![0,3],vec![3,3]]))  ;
}