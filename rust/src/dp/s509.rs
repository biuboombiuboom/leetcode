impl Solution {
    pub fn fib(n: i32) -> i32 {
        if n == 0 {
            return 0;
        }
        let mut fn0 = 0;
        let mut fn1 = 1;
        for i in 1..n {
            let num = fn0 + fn1;
            fn0 = fn1;
            fn1 = num;
        }
        return fn1;
    }
}

struct Solution;

#[test]
fn fib_test() {
    //0.1.1.2.3,5,8
    assert_eq!(Solution::fib(3), 2);
}
