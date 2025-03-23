use std::{cmp::min, collections::HashMap, i32};

impl Solution {
    pub fn min_cost(n: i32, cuts: Vec<i32>) -> i32 {
        let mut dp: HashMap<(i64, i64), i64> = HashMap::new();
        Solution::helper(0, n as i64, &cuts, &mut dp) as i32
    }

    fn helper(l: i64, r: i64, cuts: &Vec<i32>, dp: &mut HashMap<(i64, i64), i64>) -> i64 {
        if dp.contains_key(&(l, r)) {
            return *dp.get(&(l, r)).unwrap();
        }

        if l == r {
            return 0;
        }

        let i: i64 = 2;
        let mut res: i64 = i.pow(32);

        for cut in cuts.iter() {
            let i_cut: i64 = *cut as i64;
            if l < i_cut && i_cut < r {
                let curr = r - l
                    + Solution::helper(l, i_cut, cuts, dp)
                    + Solution::helper(i_cut, r, cuts, dp);
                res = min(res, curr);
            }
        }
        if res >= i.pow(32) {
            res = 0
        }
        dp.insert((l, r), res);
        return res;
    }
}
