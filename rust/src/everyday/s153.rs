pub fn find_min(nums: Vec<i32>) -> i32 {
    let mut min=nums[0];
    for i in 1..nums.len(){
        if nums[i]<nums[i-1]{
            min=nums[i];
            break
        }        
    }
    return min
}

pub fn run(){
    println!("{}",find_min(vec![3,4,1,2]))
}