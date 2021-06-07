
// left, right := 0, 0
// for _, w := range weights {
//     if w > left {
//         left = w
//     }
//     right += w
// }
// return left + sort.Search(right-left, func(x int) bool {
//     x += left
//     day := 1 // 需要运送的天数
//     sum := 0 // 当前这一天已经运送的包裹重量之和
//     for _, w := range weights {
//         if sum+w > x {
//             day++
//             sum = 0
//         }
//         sum += w
//     }
//     return day <= D
// })


impl Solution {
    pub fn ship_within_days(weights: Vec<i32>, d: i32) -> i32 {
        let m=weights.len();
        let mut left=0;
        let mut right=0;
        for w in &weights {
            left =std::cmp::max(*w,left);
            right+=w
        }
        while left<right{
            let mid=(left+right)/2;
            let mut need=1;
            let mut cur=0;
            for w in &weights{
                if cur+w>mid{
                    need+=1;
                    cur=0;
                }
                cur+=w;
            }
            if need<=d{
                right=mid;
            }else{
                left=mid+1;
            }            
        }        
        left   
    }
        
}

struct Solution {}

pub fn run() {
    println!(
        "{}",
        Solution::ship_within_days(vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 5)
    )
}

