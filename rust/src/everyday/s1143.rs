pub fn longest_common_subsequence(text1: String, text2: String) -> i32 {
    let m=text1.len();
    let n=text2.len();
    let mut dp:Vec<Vec<i32>> =Vec::new();
    for i in 0..m+1{        
        dp.push(Vec::with_capacity(n+1));
        dp[i].resize(n+1,0)
    }
    for (i,c1) in text1.chars().enumerate(){
        for (j,c2) in text2.chars().enumerate(){
            if c1==c2{
                dp[i+1][j+1] = dp[i][j] + 1; 
            }else{
                dp[i+1][j+1]=std::cmp::max(dp[i+1][j],dp[i][j+1]);               
            }
            
        }
    }
    return dp[m][n];
}

pub fn run(){
    let s1=String::from( "abc");
    let s2=String::from( "abc");
    println!("{}",longest_common_subsequence(s1,s2)) ;
}