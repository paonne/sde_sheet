use std::collections::HashMap;

impl Solution {
    pub fn can_partition(nums: Vec<i32>) -> bool {
        let target: i32 = nums.iter().sum();
        if target % 2 == 1 {
            return false;
        }
        let mut dp: HashMap<(usize, i32), bool> = HashMap::new();
        Solution::helper(&nums, 0, target / 2, &mut dp)
    }

    fn helper(
        nums: &Vec<i32>,
        ind: usize,
        target: i32,
        dp: &mut HashMap<(usize, i32), bool>,
    ) -> bool {
        if dp.contains_key(&(ind, target)) {
            return *dp.get(&(ind, target)).unwrap();
        }

        if target == 0 {
            return true;
        }

        if target < 0 || ind == nums.len() {
            return false;
        }

        let p1 = Solution::helper(nums, ind + 1, target - nums[ind], dp);
        let p2 = Solution::helper(nums, ind + 1, target, dp);

        dp.insert((ind, target), p1 || p2);
        return p1 || p2;
    }
}
