use std::collections::{HashMap, HashSet};

impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        let word_set: HashSet<String> = word_dict.into_iter().collect();
        let mut dp: HashMap<usize, bool> = HashMap::new();
        return Solution::helper(&s, &word_set, 0, &mut dp);
    }

    fn helper(s: &str, words: &HashSet<String>, ind: usize, dp: &mut HashMap<usize, bool>) -> bool {
        if dp.contains_key(&ind) {
            return *dp.get(&ind).unwrap();
        }

        if ind == s.len() {
            return true;
        }

        for i in ind + 1..s.len() + 1 {
            if words.contains(&s[ind..i]) && Solution::helper(s, words, i, dp) == true {
                dp.insert(ind, true);
                return true;
            }
        }

        dp.insert(ind, false);
        false
    }
}