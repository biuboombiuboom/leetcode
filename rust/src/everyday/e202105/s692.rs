use std::{cmp::Ordering};

struct Solution;
impl Solution {
    pub fn top_k_frequent(words: Vec<String>, k: i32) -> Vec<String> {
        let mut word_map = std::collections::HashMap::new();
        for i in 0..words.len() {
            let word=&words[i];
            let new_word = WorkdCount {
                word: word.to_string(),
                count: 0,
            };            
           let wc= word_map.entry(word).or_insert(new_word);
           wc.count+=1;
        }

        let mut list:Vec<WorkdCount>=word_map.values().cloned().collect::<Vec<WorkdCount>>();
        list.sort();
        list.reverse();    
       let list=list[0..k as usize].to_vec();
        let  ans=list.iter().map(|wc|wc.word.clone()).collect();
        return ans;
    }
}
#[derive(Clone,Eq)]
struct WorkdCount {
    word: String,
    count: i32,
}

impl Ord for WorkdCount {
    fn cmp(&self, other: &Self) -> Ordering {
        let ordering= self.count.cmp(&other.count);

        match ordering {
             Ordering::Equal=>self.word.cmp(&other.word).reverse(),
             Ordering::Greater=>ordering,
             Ordering::Less=>ordering,
        }
       
    }
}

impl PartialOrd for WorkdCount {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl PartialEq for WorkdCount {
    fn eq(&self, other: &Self) -> bool {
        self.count == other.count
    }
}

pub fn run() {
    println!("{:?}", Solution::top_k_frequent(vec!["i", "love", "leetcode", "i", "love", "coding"].iter().map(|a|a.to_string()).collect(), 3));
}
