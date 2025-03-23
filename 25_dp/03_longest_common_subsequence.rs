use std::{cmp::max, collections::HashMap, str};

impl Solution {
    pub fn longest_common_subsequence(text1: String, text2: String) -> i32 {
        let mut dp: HashMap<(usize, usize), i32> = HashMap::new();
        Solution::helper(&text1, &text2, 0, 0, &mut dp)
    }

    fn helper(
        text1: &str,
        text2: &str,
        ind1: usize,
        ind2: usize,
        dp: &mut HashMap<(usize, usize), i32>,
    ) -> i32 {
        if dp.contains_key(&(ind1, ind2)) {
            return *dp.get(&(ind1, ind2)).unwrap();
        }

        if ind1 == text1.len() || ind2 == text2.len() {
            return 0;
        }

        if text1.as_bytes()[ind1] == text2.as_bytes()[ind2] {
            let res = 1 + Solution::helper(text1, text2, ind1 + 1, ind2 + 1, dp);
            dp.insert((ind1, ind2), res);
            return res;
        }

        let p1 = Solution::helper(text1, text2, ind1 + 1, ind2, dp);
        let p2 = Solution::helper(text1, text2, ind1, ind2 + 1, dp);

        dp.insert((ind1, ind2), max(p1, p2));
        return *dp.get(&(ind1, ind2)).unwrap();
    }
}
