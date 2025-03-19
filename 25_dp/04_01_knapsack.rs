use std::{cmp::max, collections::HashMap};

struct Solution {}

impl Solution {
    fn zero_one_knapsack(wt: Vec<usize>, val: Vec<usize>, w: usize) -> usize {
        let mut dp: HashMap<(usize, usize), usize> = HashMap::new();
        Solution::helper(&wt, &val, 0, w, &mut dp)
    }

    fn helper(
        wt: &Vec<usize>,
        val: &Vec<usize>,
        ind: usize,
        curr: usize,
        dp: &mut HashMap<(usize, usize), usize>,
    ) -> usize {
        if dp.contains_key(&(ind, curr)) {
            return *dp.get(&(ind, curr)).unwrap();
        }

        if ind == wt.len() {
            return 0;
        }

        let p1 = Solution::helper(wt, val, ind + 1, curr, dp);

        let mut p2 = 0;
        if wt[ind] <= curr {
            p2 = val[ind] + Solution::helper(wt, val, ind + 1, curr - wt[ind], dp);
        }

        dp.insert((ind, curr), max(p1, p2));
        *dp.get(&(ind, curr)).unwrap()
    }
}

fn main() {
    assert_eq!(
        Solution::zero_one_knapsack(vec![10, 20, 30], vec![60, 100, 120], 50),
        220
    );
    assert_eq!(
        Solution::zero_one_knapsack(vec![5, 4, 6, 3], vec![10, 40, 30, 50], 10),
        90
    );
    assert_eq!(
        Solution::zero_one_knapsack(vec![1, 2, 3, 8, 7, 4], vec![20, 5, 10, 40, 15, 25], 10),
        60
    );
}
