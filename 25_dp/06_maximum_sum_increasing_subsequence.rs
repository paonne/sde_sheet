use std::cmp::max;

struct Solution {}

impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut dp: Vec<i32> = nums.clone();

        for i in (0..n).rev() {
            for j in i + 1..n {
                if nums[i] < nums[j] {
                    dp[i] = max(dp[i], nums[i] + dp[j]);
                }
            }
        }

        *dp.iter().max().unwrap()
    }
}

fn main() {
    assert_eq!(Solution::length_of_lis(vec![1, 101, 2, 3, 100]), 106);
    assert_eq!(Solution::length_of_lis(vec![4, 1, 2, 3]), 6);
    assert_eq!(Solution::length_of_lis(vec![4, 1, 2, 4]), 7);
}
