// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::{rc::Rc};
use std::{cell::RefCell, usize};
struct Solution;
impl Solution {
    pub fn is_cousins(root: Option<Rc<RefCell<TreeNode>>>, x: i32, y: i32) -> bool {
        if x == y {
            return false;
        }
        let mut nodes: Vec<Option<Rc<RefCell<TreeNode>>>> = Vec::new();
        nodes.push(root);
        loop {
            let mut px = usize::MAX;
            let mut py = usize::MAX;
            let mut childs: Vec<Option<Rc<RefCell<TreeNode>>>> = Vec::new();
            for i in 0..nodes.len() {
                if let Some(node) = nodes[i].clone() {
                    if let Some(left) = node.borrow().left.clone() {
                        childs.push(node.borrow().left.clone());
                        if left.borrow().val == x {
                            px = i;
                        }
                        if left.borrow().val == y {
                            py = i;
                        }
                    }
                    if let Some(right) = node.borrow().right.clone() {
                        childs.push(node.borrow().right.clone());
                        if right.borrow().val == x {
                            px = i;
                        }
                        if right.borrow().val == y {
                            py = i;
                        }
                    }
                }
                if px == py && py != usize::MAX {
                    return false;
                }
            }
            if px!=usize::MAX&&py!=usize::MAX{
                return true
            }

            if childs.len() == 0 {
                break;
            }
            nodes = childs;
        }
        false
    }
}

pub fn run() {
    

    let root = Rc::new(RefCell::new(TreeNode::new(1)));
    let left=Rc::new(RefCell::new(TreeNode::new(2)));
    let right=Rc::new(RefCell::new(TreeNode::new(3)));
    left.borrow_mut().left=Some(Rc::new(RefCell::new(TreeNode::new(4))));
    root.borrow_mut().left=Some(left);
    root.borrow_mut().right=Some(right);

   
    let is_cousins = Solution::is_cousins(Some(root),3, 4, );
    println!("{}", is_cousins);
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
