fn largest_number(nums: Vec<i32>) -> String {
    let mut new_nums: Vec<i32> = Vec::new();
    let mut zero_count=0;
    for num in nums.iter() {
        let mut find = false;
        if *num==0{
            zero_count+=1;
        }
        for i in 0..new_nums.len() {
            if cmp_bg(*num, new_nums[i]){   
                find = true;              
                new_nums.insert(i, *num);               
                break;
            }
        }
        if !find  {
            new_nums.push(*num);
        }
    }    
    if zero_count==nums.len(){
        return String::from("0")
    }

    return new_nums
        .into_iter()
        .map(|d| d.to_string())
        .collect::<String>();
}

fn cmp_bg(x: i32, y: i32) -> bool {
    if x==y{
        return true
    }
    let xs=x.to_string();
    let ys=y.to_string();
    let lx = xs.len();
    let ly = ys.len();
    let x_bytes =xs.as_bytes();
    let y_bytes = ys.as_bytes();
    let mut i = 0;
    loop {
        let x1 = x_bytes[i % lx];
        let y1 = y_bytes[i % ly];

        if x1 != y1 {
            return x1 > y1;
        }        
        i += 1;
        if i%lx==0&&i%ly==0{
            return true
        }
    }
}

pub fn run() {
    println!("{}", largest_number(vec![3,43,48,94,85,33,64,32,63,66]));
}

// 0 1
