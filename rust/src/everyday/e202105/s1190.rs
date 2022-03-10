struct Solution;
impl Solution {
    pub fn reverse_parentheses(s: String) -> String {
        let mut stack: Vec<u8> = Vec::new();
        for &b in s.as_bytes().iter() {
            if b == b')' {
                let mut reverse_stack: Vec<u8> = Vec::new();
                while let Some(top) = stack.pop() {
                    if top == b'(' {
                        break;
                    }
                    reverse_stack.push(top);
                }
                stack.append(&mut reverse_stack);
                continue;
            }
            stack.push(b);
        }
        return String::from_utf8(stack).unwrap();
    }
}

#[test]
pub fn reverse_parentheses_test() {
    assert_eq!("apmnolkjihgfedcbq", Solution::reverse_parentheses(String::from("a(bcdefghijkl(mno)p)q")));
}
