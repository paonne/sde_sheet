use std::cmp::max;

impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut lis: Vec<i32> = vec![1; n];

        for i in (0..n).rev() {
            for j in i + 1..n {
                if nums[i] < nums[j] {
                    lis[i] = max(lis[i], 1 + lis[j]);
                }
            }
        }

        *lis.iter().max().unwrap()
    }
}
