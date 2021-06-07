pub fn search(nums: Vec<i32>, target: i32) -> bool {
    for i in 0..nums.len(){
        if nums[i]==target{
            return true
        }
        
        if i>0&& nums[i]<nums[i-1] {            
            if nums[i-1]<target||target<nums[i]{
                println!("xxx{}",i);
                return false
            }
        }
    }

    return false
}

pub fn run(){
    println!("{}",search(vec![1],1))
}