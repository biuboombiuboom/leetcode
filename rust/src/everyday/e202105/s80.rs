pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
    let mut i=2;
    while i <nums.len() {               
        if nums[i]==nums[i-1]&&nums[i]==nums[i-2]{
            nums.remove(i);
            i-=1;
        }
        i+=1
    }
    return nums.len() as i32
}


pub fn run(){
    let mut nums:Vec<i32>=vec![1,1,1,1,2,2,3];
    remove_duplicates(&mut nums);
   
        println!("{:?}",nums);
}