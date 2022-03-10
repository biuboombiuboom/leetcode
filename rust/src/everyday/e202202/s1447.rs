impl Solution {
    pub fn simplified_fractions(n: i32) -> Vec<String> {
        let mut ans = Vec::new();
        for i in 1..n {
            for j in i + 1..n + 1 {
                let gcd = gcd(j, i);
                if gcd == 1 {
                    ans.push(format!("{}/{}", i, j));
                }
            }
        }
        return ans;
    }
}

fn gcd(a: i32, b: i32) -> i32 {
    if a % b == 0 {
        return b;
    }
    return gcd(b, a % b);
}

struct Solution;

#[test]
fn simplified_fractions_test() {
    let input = 8;
    let output = vec![
        "1/2", "1/3", "1/4", "1/5", "1/6", "1/7", "1/8", "2/3", "2/5", "2/7", "3/4", "3/5", "3/7",
        "3/8", "4/5", "4/7", "5/6", "5/7", "5/8", "6/7", "7/8",
    ];
    assert_eq!(output, Solution::simplified_fractions(input));
}
