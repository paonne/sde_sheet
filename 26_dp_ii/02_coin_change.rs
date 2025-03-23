use std::{cmp::min, collections::HashMap};

impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let mut dp: HashMap<(usize, i32), i32> = HashMap::new();
        let res = Solution::helper(&coins, 0, 0, amount, &mut dp);
        if res == i32::MAX {
            return -1;
        }
        return res;
    }

    fn helper(
        coins: &Vec<i32>,
        ind: usize,
        curr: i32,
        amount: i32,
        dp: &mut HashMap<(usize, i32), i32>,
    ) -> i32 {
        if dp.contains_key(&(ind, curr)) {
            return *dp.get(&(ind, curr)).unwrap();
        }
        if ind == coins.len() || curr > amount {
            return i32::MAX;
        }

        if curr == amount {
            return 0;
        }

        let p1: i64 = 1 + Solution::helper(coins, ind, curr + coins[ind], amount, dp) as i64;
        let p2: i64 = Solution::helper(coins, ind + 1, curr, amount, dp) as i64;
        let mini = min(p1, p2);
        if mini >= i32::MAX as i64 {
            dp.insert((ind, curr), i32::MAX);
            return i32::MAX;
        }
        dp.insert((ind, curr), mini as i32);
        return mini as i32;
    }
}
