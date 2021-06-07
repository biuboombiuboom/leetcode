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
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn range_sum_bst(root: Option<Rc<RefCell<TreeNode>>>, low: i32, high: i32) -> i32 {
        let mut sum=0;
        fn search(root: Option<Rc<RefCell<TreeNode>>>,sum:&mut i32,low:i32,high:i32){
            if let Some(node)=root{
                let val=node.borrow().val;
                if val<low{
                    search(node.borrow().right.clone(), sum, low, high);
                }
                if val>high{
                    search(node.borrow().left.clone(), sum, low, high);
                }else{
                    search(node.borrow().left.clone(), sum, low, high);
                    *sum+=val;
                    search(node.borrow().right.clone(), sum, low, high);
                   
                }
            }
        }
        search(root, &mut sum, low, high);
        return sum;
        
        //     let ans = Rc::new(RefCell::new(0));
        //     struct Search<'s> {
        //         f: &'s dyn Fn(&Search, Option<Rc<RefCell<TreeNode>>>, Rc<RefCell<i32>>),
        //     }

        //     let search=Search{
        //         f:&|search,node,ans|{
        //             if let Some(node)=node{
        //                 if node.borrow().val<low||node.borrow().val>high{
        //                     return
        //                 }
        //             }
        //         },
        //     };

        //    (search.f)(&search, root,ans.clone());

        //     return *ans.borrow();

       
        // let mut ans = 0;
        // if let Some(node) = root {
        //     if node.borrow().val >= low && node.borrow().val <= high {
        //         ans += node.borrow().val;
        //     }
        //     ans += Solution::range_sum_bst(node.borrow().left.clone(), low, high);
        //     ans += Solution::range_sum_bst(node.borrow().right.clone(), low, high);
        // }
        // ans
    }
}

struct Solution {}

pub fn run() {
    let root = Some(Rc::new(RefCell::new(TreeNode::new(10))));
    println!("{}", Solution::range_sum_bst(root, 7, 15));
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

fn build(vals: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
    for val in vals {
        let node = TreeNode::new(val);
    }
    None
}
