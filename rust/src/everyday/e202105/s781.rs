use std::collections::HashMap;

pub fn num_rabbits(answers: Vec<i32>) -> i32 {
    let mut map:HashMap<i32,i32> = HashMap::new(); 
    if answers.len()==0{
        return 0
    } 
    let mut min_count=0;
    for i in 0..answers.len(){        
        let r=answers[i];
        match map.get(&r){
            Some(c)=>{
                if *c==r+1{
                    min_count+=1;
                    min_count+=r;
                    map.insert(r,1);
                }else{
                    map.insert(r,*c+1);
                }
            }
            None=>{
                min_count+=1;
                min_count+=r;
                map.insert(r,1);
            },
        }
    }
    return min_count

}

pub fn run(){
    let count=num_rabbits(vec![0,0,1,1,1]);
    println!("{}",count);
}