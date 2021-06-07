use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn min_diff_in_bst(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let ans = Rc::new(RefCell::new(106));
        let pre = Rc::new(RefCell::new(-1));
        struct Dfs<'s> {
            f: &'s dyn Fn(&Dfs, Option<Rc<RefCell<TreeNode>>>, Rc<RefCell<i32>>, Rc<RefCell<i32>>),
        }
        let dfs = Dfs {
            f: &|dfs, node, ans, pre| {
                if let Some(n) = node {
                    if let Some(left) = n.borrow().left.clone() {
                        (dfs.f)(dfs, Some(left), ans.clone(), pre.clone());
                    }
                    let diff = n.borrow().val - *pre.borrow();
                    if *(pre.borrow()) != -1 && diff < *ans.borrow() {
                        *ans.borrow_mut() = diff;
                    }
                    *pre.borrow_mut() = n.borrow().val;
                    if let Some(right) = n.borrow().right.clone() {
                        (dfs.f)(dfs, Some(right), ans.clone(), pre.clone());
                    }
                }
            },
        };

        (dfs.f)(&dfs, root, ans.clone(), pre);
        return *ans.borrow();
    }
}

struct Solution {}
pub fn run() {
    let root = TreeNode::new(1);
    let rc = Rc::new(RefCell::new(root));
    println!("{}", Solution::min_diff_in_bst(Some(rc)));
}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

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
