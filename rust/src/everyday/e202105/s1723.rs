use core::time;
use std::{collections::{HashMap, HashSet}, usize};


struct Solution{}
impl Solution {
    pub fn minimum_time_required(jobs: Vec<i32>, k: i32) -> i32 {
        let mut jobs=jobs;
        jobs.sort();
        let mut job_map:HashMap<usize,Vec<i32>>=HashMap::new();
        let mut time_map:HashMap<usize,i32>=HashMap::new();
        let mut total=0;
        let n =jobs.len();       
        for i in 0..n {
            let job=jobs[i];         
            let key=i%k as usize;

            let sub_jobs=job_map.entry(key).or_insert(vec![]);
            sub_jobs.push(job);
            let time=time_map.entry(key).or_insert(0);
            *time+=job;
            total+=job;
        }
        let avg=total/k;
        
        0
    }
}


pub fn run(){    
    assert_eq!(11, Solution::minimum_time_required(vec![1,2,4,7,8],2));
}

