struct Solution;
impl Solution {
    pub fn num_decodings(s: String) -> i32 {
        let m = 1000000000 + 7;
        let mut a = 0;
        let mut b = 1;
        let mut c = 0;
        for (i, ch) in s.bytes().enumerate() {
            c = b * Solution::check1digit(ch) % m;
            if i > 0 {
                c = (c + a * Solution::check2digits(s.as_bytes()[i - 1], ch)) % m;
            }
            a = b;
            b = c;
        }

        return c as i32;
    }

    fn check1digit(c: u8) -> i64 {
        if c == b'0' {
            return 0;
        }
        if c == b'*' {
            return 9;
        }
        return 1;
    }

    fn check2digits(c0: u8, c1: u8) -> i64 {
        if c0 == b'*' && c1 == b'*' {
            return 15;
        }
        if c0 == b'*' {
            if c1 <= b'6' {
                return 2;
            }
            return 1;
        }

        if c1 == b'*' {
            if c0 == b'1' {
                return 9;
            }
            if c0 == b'2' {
                return 6;
            }
            return 0;
        }
        if c0 != b'0' && (c0 - b'0') * 10 + (c1 - b'0') <= 26 {
            return 1;
        }

        return 0;
    }
}

#[test]
fn num_decodings_test() {
    assert_eq!(
        Solution::num_decodings("7*9*3*6*3*0*5*4*9*7*3*7*1*8*3*2*0*0*6*".to_string()),
        18
    );
}
