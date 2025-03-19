use std::cmp::max;

impl Solution {
    pub fn max_product(nums: Vec<i32>) -> i32 {
        let (mut pre_max, mut suff_max) = (1, 1);
        let n = nums.len();
        let mut res = nums[0];

        for i in 0..n {
            pre_max = pre_max * nums[i];
            suff_max = suff_max * nums[n - i - 1];

            res = max(res, max(pre_max, suff_max));

            if pre_max == 0 {
                pre_max = 1
            }
            if suff_max == 0 {
                suff_max = 1
            }
        }

        res
    }
}
