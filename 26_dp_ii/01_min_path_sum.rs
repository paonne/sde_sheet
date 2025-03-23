use std::{cmp::min, collections::HashMap};

impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let mut dp: HashMap<(usize, usize), i32> = HashMap::new();
        return Solution::helper(&grid, 0, 0, &mut dp);
    }

    fn helper(
        grid: &Vec<Vec<i32>>,
        r: usize,
        c: usize,
        dp: &mut HashMap<(usize, usize), i32>,
    ) -> i32 {
        if dp.contains_key(&(r, c)) {
            return *dp.get(&(r, c)).unwrap();
        }
        if r == grid.len() - 1 && c == grid[0].len() - 1 {
            return grid[r][c];
        }

        if r == grid.len() || c == grid[0].len() {
            return i32::MAX;
        }

        let p1 = Solution::helper(grid, r + 1, c, dp);
        let p2 = Solution::helper(grid, r, c + 1, dp);
        let res = grid[r][c] + min(p1, p2);

        dp.insert((r, c), res);
        return res;
    }
}
