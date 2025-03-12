use std::collections::HashSet;

impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> Vec<String> {
        let mut res: Vec<String> = Vec::new();
        let mut word_set: HashSet<String> = word_dict.into_iter().collect();
        let mut curr: Vec<String> = Vec::new();
        Solution::helper(&s, &word_set, &mut res, 0, &mut curr);
        res
    }

    fn helper(
        s: &String,
        word_set: &HashSet<String>,
        res: &mut Vec<String>,
        ind: usize,
        curr: &mut Vec<String>,
    ) {
        if ind == s.len() {
            res.push(curr.join(" "));
            return;
        }

        for i in ind..s.len() {
            let curr_string = &s[ind..i + 1];
            if word_set.contains(curr_string) {
                curr.push(curr_string.to_string());
                Solution::helper(&s, word_set, res, i + 1, curr);
                curr.pop();
            }
        }
    }
}