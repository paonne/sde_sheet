use std::{cmp::min, collections::HashMap};

impl Solution {
    pub fn min_distance(word1: String, word2: String) -> i32 {
        let mut dp: HashMap<(usize, usize), usize> = HashMap::new();
        Solution::helper(&word1, &word2, 0, 0, &mut dp) as i32
    }

    fn helper(
        word1: &str,
        word2: &str,
        ind1: usize,
        ind2: usize,
        dp: &mut HashMap<(usize, usize), usize>,
    ) -> usize {
        if dp.contains_key(&(ind1, ind2)) {
            return *dp.get(&(ind1, ind2)).unwrap();
        }

        if ind1 == word1.len() && ind2 == word2.len() {
            return 0;
        }

        if ind1 == word1.len() {
            return word2.len() - ind2;
        }

        if ind2 == word2.len() {
            return word1.len() - ind1;
        }

        if word1.as_bytes()[ind1] == word2.as_bytes()[ind2] {
            let res = Solution::helper(word1, word2, ind1 + 1, ind2 + 1, dp);
            dp.insert((ind1, ind2), res);
            return res;
        }

        let insert = Solution::helper(word1, word2, ind1, ind2 + 1, dp);
        let delete = Solution::helper(word1, word2, ind1 + 1, ind2, dp);
        let replace = Solution::helper(word1, word2, ind1 + 1, ind2 + 1, dp);

        let res = 1 + min(insert, min(delete, replace));
        dp.insert((ind1, ind2), res);
        res
    }
}
