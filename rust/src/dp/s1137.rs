impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        if n == 0 {
            return 0;
        }
        if n <= 2 {
            return 1;
        }
        let mut fn0 = 0;
        let mut fn1 = 1;
        let mut fn2 = 1;

        for _ in 3..n + 1 {
            let num = fn0 + fn1 + fn2;
            fn0 = fn1;
            fn1 = fn2;
            fn2 = num;
        }

        return fn2;
    }
}

struct Solution;

#[test]
fn tribonacci_test() {
    //0,1,1,2,4,7
    assert_eq!(Solution::tribonacci(3), 2)
}
