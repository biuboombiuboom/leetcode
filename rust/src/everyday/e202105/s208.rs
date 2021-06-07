use std::cell::RefCell;
use std::rc::Rc;
struct Trie {
    key: u8,
    inserted: bool,
    childs: Rc<RefCell<Vec<Trie>>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Trie {
    /** Initialize your data structure here. */
    fn new() -> Self {
        return Self {
            key: 0,
            inserted: false,
            childs: Rc::new(RefCell::new(Vec::new())),
        };
    }
    /** Inserts a word into the trie. */
    fn insert(&self, word: String) {
        if word.len() == 0 {
            return;
        }
        let index = (word.as_bytes()[0] - b'a') as usize;
        let c_len = self.childs.borrow().len();
        for j in c_len..index + 1 {
            let mut new_node = Trie::new();
            new_node.key = b'a' + j as u8;
            self.childs.borrow_mut().push(new_node);
        }
        self.childs.borrow_mut()[index].inserted = true;
        self.childs.borrow()[index].insert(word[1..].to_string());
    }
    /** Returns if the word is in the trie. */
    fn search(&self, word: String) -> bool {        
        let index = (word.as_bytes()[0] - b'a') as usize;
        let c_len = self.childs.borrow().len();
        if c_len <= index || !self.childs.borrow_mut()[index].inserted {
            return false;
        }
        if word.len()==1{
            return true
        }
        return self.childs.borrow()[index].search(word[1..].to_string());
    }
    /** Returns if there is any word in the trie that starts with the given prefix. */
    fn starts_with(&self, prefix: String) -> bool {
        let index = (prefix.as_bytes()[0] - b'a') as usize;
        let c_len = self.childs.borrow().len();
        if c_len <= index || !self.childs.borrow_mut()[index].inserted {
            return false;
        }
        if prefix.len()==1{
            return true
        }
        return self.childs.borrow()[index].search(prefix[1..].to_string());
    }
}

pub fn run() {
    let t = Trie::new();
    t.insert(String::from("apple"));
    println!("{}", t.starts_with(String::from("abp")));
}
