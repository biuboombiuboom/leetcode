pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
    let mut i = m - 1;
    let mut j = n - 1;
    let mut k = (m + n ) as usize;

    loop {
        let mut num1: i32 = -110;
        let mut num2: i32 = -110;
        if i >= 0 {
            num1 = nums1[i as usize];
        }
        if j >=0 {
            num2 = nums2[j as usize ];
        }
        if num1 > num2 {
            nums1[k-1] = num1;
            i -= 1;
        } else {
            nums1[k-1] = num2;
            j -= 1;
        }  
        k-=1;
        if k == 0 {
            break;
        }
    }
    if j>=0{
        for m in 0..(j+1) as usize{
            nums1[m]=nums2[m];
        }
    }
}

pub fn run() {
    let mut nums1 = vec![1];
    let mut nums2 = vec![0];
    merge(&mut nums1, 1, &mut nums2, 0);
    println!("{:?}", nums1);
}
