struct Solution;
impl Solution {
    pub fn convert_to_base7(num: i32) -> String {
        if num == 0 {
            return String::from("0");
        }
        let mut num = num;
        let mut ans = String::default();
        let mut flag = false;
        if num < 0 {
            num = num * -1;
            flag = true;
        }
        while num >= 7 {
            let b = num % 7;
            num = num / 7;
            ans.insert_str(0, b.to_string().as_str())
        }
        if num != 0 {
            ans.insert_str(0, num.to_string().as_str())
        }
        if flag {
            ans.insert(0, '-');
        }
        return ans;
    }
}

#[test]
fn convert_to_base7_test() {
    assert_eq!(Solution::convert_to_base7(-1), "-1");
}
