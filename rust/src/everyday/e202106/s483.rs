struct Solution;
impl Solution {
    pub fn smallest_good_base(n: String) -> String {
        "".to_string()
    }
}

#[test]
fn smallest_good_base_test(){
    assert_eq!(Solution::smallest_good_base("".to_string()),"".to_string());
}