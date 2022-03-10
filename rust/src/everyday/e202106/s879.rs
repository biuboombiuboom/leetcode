struct Solution;
impl Solution {
    pub fn profitable_schemes(n: i32, min_profit: i32, group: Vec<i32>, profit: Vec<i32>) -> i32 {
        let base:i32=10;
        let mmod=base.pow(9)+7;
        let m=group.len();
        let n=n as usize;
        let min_profit=min_profit as usize;
        let mut dp=Vec::new();
        for i in 0..m+1{
            dp.push(Vec::new());
            for j in 0..n+1{
                dp[i].push(Vec::new());
                dp[i][j].resize(min_profit+1 ,0);
            }
        }
        dp[0][0][0]=1;
        for i in 0..m{
            let p=profit[i] as usize;
            let g=group[i] as usize;
            for j in 0..n+1{
                for k in 0..min_profit+1{
                    if j<g{
                        dp[i+1][j][k]= dp[i][j][k];
                    }else{
                        let jg=j-g;
                        let mut kk=0;
                        if k>p{
                            kk=k-p;
                        }
                        let ddd=dp[i][jg][kk];
                        dp[i+1][j][k] = (dp[i][j][k] + ddd) % mmod;
                    }
                }
            }
        }
        let mut sum=0;
        for d in dp[m].iter() {
            sum = (sum + d[min_profit]) % mmod
        }
        return sum;
    
    }
}

#[test]
fn profitable_schemes_test(){
    assert_eq!(Solution::profitable_schemes(10,5,vec![2,3,5],vec![6,7,8]),7);
}