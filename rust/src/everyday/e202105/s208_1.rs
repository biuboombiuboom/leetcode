#[derive(Default)]
struct Trie {
    node: CharNode,
}

#[derive(Default)]
struct CharNode {
    children: [Option<Box<CharNode>>; 26],
    is_world: bool,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Trie {
    /** Initialize your data structure here. */
    fn new() -> Self {
        Self::default()
    }

    /** Inserts a word into the trie. */
    fn insert(&mut self, word: String) {
        let mut node = &mut self.node;
        if word.len() == 0 {
            return;
        }
        for b in word.as_bytes() {
            node = node.children[(*b - b'a') as usize]
                .get_or_insert_with(|| Box::new(CharNode::default()));
        }
        node.is_world = true
    }

    /** Returns if the word is in the trie. */
    fn search(&mut self, word: String) -> bool {
        let mut node = &mut self.node;
        for b in word.as_bytes() {
            if let Some(ref mut t) = node.children[(*b - b'a') as usize] {
                node = t;
            } else {
                return false;
            }
        }
        return node.is_world;
    }

    /** Returns if there is any word in the trie that starts with the given prefix. */
    fn starts_with(&mut self, prefix: String) -> bool {
        let mut node = &mut self.node;
        for b in prefix.as_bytes() {
            if let Some(ref mut t) = node.children[(*b - b'a') as usize] {
                node = t;
            } else {
                return false;
            }
        }
        return true;
    }
}
