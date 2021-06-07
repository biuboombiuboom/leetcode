pub fn trap(height: Vec<i32>) -> i32 {    
   let mut highest=0;
   let n=height.len();
    for i in 0..height.len() {
        if height[i]>height[highest]{
            highest=i;
        }
    }
    if n<3{
        return 0
    }
    let mut cap=0;
    let mut l=height[0];
    for i in 1..highest{
        if height[i]<l{
            cap+=l-height[i]
        }else{
            l=height[i]
        }
    }

    l=height[n-1];
    for i in (highest+1..n-1).rev() {
        if height[i] < l {
			cap += l - height[i]
		} else {
			l = height[i]
		}  
    }

    return cap;
}

pub fn run(){
   let cap= trap(vec![2,0,2]);
   println!("{}",cap);
}