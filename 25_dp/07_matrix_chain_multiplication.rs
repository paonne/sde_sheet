use std::{cmp::min, collections::HashMap};

struct Solution {}

impl Solution {
    fn matrix_chain_multiplication(nums: Vec<usize>) -> usize {
        let mut dp: HashMap<(usize, usize), usize> = HashMap::new();
        Solution::helper(&nums, 1, nums.len() - 1, &mut dp)
    }

    fn helper(
        nums: &Vec<usize>,
        i: usize,
        j: usize,
        dp: &mut HashMap<(usize, usize), usize>,
    ) -> usize {
        if dp.contains_key(&(i, j)) {
            return *dp.get(&(i, j)).unwrap();
        }

        if i == j {
            return 0;
        }

        let mut res = usize::MAX;

        for k in i..j {
            res = min(
                res,
                Solution::helper(nums, i, k, dp)
                    + Solution::helper(nums, k + 1, j, dp)
                    + (nums[i - 1] * nums[k] * nums[j]),
            );
        }

        dp.insert((i, j), res);
        return *dp.get(&(i, j)).unwrap();
    }
}

fn main() {
    assert_eq!(Solution::matrix_chain_multiplication(vec![2, 1, 3, 4]), 20);
    assert_eq!(
        Solution::matrix_chain_multiplication(vec![1, 2, 3, 4, 3]),
        30
    );
    assert_eq!(Solution::matrix_chain_multiplication(vec![3, 4]), 0)
}
