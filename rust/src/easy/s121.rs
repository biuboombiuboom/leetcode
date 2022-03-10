pub fn max_profit(prices: Vec<i32>) -> i32 {   
    let mut dp:Vec<i32>=Vec::new();
    let mut min_price=prices[0];
    dp.push(0);
    for i in 1..prices.len(){
        if min_price>prices[i]{
            min_price=prices[i]
        }
        dp.push(max(dp[i-1],prices[i]-min_price))
    }
    
    return dp[dp.len()-1]
}

fn max(x: i32, y: i32) -> i32 {
    if x > y {
        return x;
    }
    return y;
}

pub fn run() {
    let max=max_profit(vec![7,1,5,3,6,4]);
    println!("{}", max);
}
