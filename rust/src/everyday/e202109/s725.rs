#[derive(PartialEq, Eq, Clone, Debug)]
struct ListNode {
    val: i32,
    next: Option<Box<ListNode>>,
}

impl ListNode {
    fn new(val: i32) -> Self {
        return Self {
            val: val,
            next: None,
        };
    }

    fn append(&mut self, num: i32) {
        let new_node = Some(Box::new(ListNode::new(num)));
        let tail = &mut self.next;
        match tail {
            Some(ref mut tail) => tail.append(num),
            None => self.next = new_node,
        }
    }
}

struct Solution;
impl Solution {
    pub fn split_list_to_parts(head: Option<Box<ListNode>>, k: i32) -> Vec<Option<Box<ListNode>>> {
        let mut ans = vec![None; k as usize];
        let mut len = 0;

        let mut iter_node = &head;

        while let Some(node) = iter_node {
            len += 1;
            iter_node = &node.next
        }
        let part_len = len / k;
        let rem = len - k * part_len;
        let mut curr = head.as_ref();
        for i in 0..k {
            let mut part_len0 = part_len;
            if i < rem {
                part_len0 += 1;
            }
            let mut begin = Some(Box::new(ListNode::new(0)));
            let mut tail = begin.as_mut();
            for _ in 0..part_len0 {
                if let Some(node) = curr {
                    if let Some(mut node1) = tail {
                        node1.next = Some(Box::new(ListNode::new(node.val)));
                        tail = node1.next.as_mut();
                    }
                    curr = node.next.as_ref();
                }
            }
            ans[i as usize] = begin.unwrap().next;
        }
        return ans;
    }
}

#[test]
fn split_list_to_parts_test() {
    let mut begin = ListNode::new(0);
    vec![1, 2, 3].iter().for_each(|x| begin.append(*x));
    println!("{:?}", begin);
    let head: Option<Box<ListNode>> = begin.next;
    let k = 5;
    let mut ans: Vec<Option<Box<ListNode>>> = vec![vec![1], vec![2], vec![3], vec![], vec![]]
        .iter()
        .map(|item| {
            let mut begin0 = ListNode::new(0);
            item.iter().for_each(|x| begin0.append(*x));
            return begin0.next;
        })
        .collect();
    println!("{:?}", head);
    let ans1 = Solution::split_list_to_parts(head, k);
    println!("{:?}", ans1);
    println!("{:?}", ans);
    assert_eq!(ans1, ans)
}
