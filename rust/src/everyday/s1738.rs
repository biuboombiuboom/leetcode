
        use std::cmp::Reverse;

struct Solution;
impl Solution {
    pub fn kth_largest_value(matrix: Vec<Vec<i32>>, k: i32) -> i32 {
        let k =k as usize;
        let mut heap=std::collections::BinaryHeap::with_capacity(k);
        let mut new_matrix:Vec<Vec<i32>>=Vec::new();
        let m=matrix.len();
        let n=matrix[0].len();
        let mut top_k:Vec<i32>=Vec::new();
        for i in 0..m{
            new_matrix.push(Vec::new());
            for j in 0..n{    
                new_matrix[i].push(matrix[i][j])            ;
                if i>0{
                    new_matrix[i][j]= new_matrix[i][j]^new_matrix[i-1][j];
                }
                if j>0{
                    new_matrix[i][j]= new_matrix[i][j]^new_matrix[i][j-1];
                }
                if i>0&&j>0{
                    new_matrix[i][j]= new_matrix[i][j]^new_matrix[i-1][j-1];
                }   
                if heap.len()<k{
                    heap.push(Reverse(new_matrix[i][j]))
                }else if heap.peek().unwrap().0<new_matrix[i][j]{
                    heap.pop();
                    heap.push(Reverse(new_matrix[i][j]));
                }
            }            
        }
        return heap.peek().unwrap().0;       
    }
}




pub fn run(){
    println!("{}",Solution::kth_largest_value(vec![vec![5,2],vec![1,6]], 2))
}