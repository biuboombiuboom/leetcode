struct Solution;

use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}
//
impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

impl Solution {
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> i32 {
        let mut ans = 0;
        if let Some(node) = root.clone() {
            ans = Solution::rootSum(root, target_sum);
            ans += Solution::path_sum(node.borrow().left.clone(), target_sum);
            ans += Solution::path_sum(node.borrow().right.clone(), target_sum);
        }
        return ans;
    }

    fn rootSum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> i32 {
        let mut ans = 0;
        if let Some(node) = root {
            let val = node.borrow().val;
            if val == target_sum {
                ans += 1;
            }
            ans += Solution::rootSum(node.borrow().left.clone(), target_sum - val);
            ans += Solution::rootSum(node.borrow().right.clone(), target_sum - val);
        }
        return ans;
    }
}
