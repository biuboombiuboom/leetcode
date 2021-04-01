pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut nums = nums;
    nums.sort();
    let mut ans: Vec<Vec<i32>> = Vec::new();
    let n = nums.len();
    let mut first = 0;
    while first < n {
        if first > 0 && nums[first] == nums[first - 1] {
            first += 1;
            continue;
        }
        let mut third = n - 1;
        let mut second = first + 1;
        let target = -1 * nums[first];

        while second < n {
            if second > first + 1 && nums[second] == nums[second - 1] {
                second += 1;
                continue;
            }
            while second < third && nums[second] + nums[third] > target {
                third -= 1;
            }
            if second == third {
                break;
            }
            if nums[second] + nums[third] == target {
                ans.push(vec![nums[first], nums[second], nums[third]]);
            }
            second += 1;
        }

        first += 1;
    }
    return ans;
}

pub fn run() {
    let output = three_sum(vec![1, 0, -1, 2, -1]);
    for elem in output {
        println!("{},{},{}", elem[0], elem[1], elem[2]);
    }
}
