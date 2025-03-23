use std::{cmp::min, collections::HashMap};

fn palindrom_partition(s: String) -> usize{
    let mut dp: HashMap<usize, usize> = HashMap::new();
    return helper(&s, 0, &mut dp)
}

fn helper(s: &str, ind: usize, dp: &mut HashMap<usize, usize>) -> usize{
    if dp.contains_key(&ind){
        return *dp.get(&ind).unwrap()
    }
    if &s[ind..] == &s[ind..].chars().rev().collect::<String>(){
        return 0
    }
    let mut res = s.len();
    
    for i in ind+1..s.len(){
        let curr_string = &s[ind..i];
        if curr_string == curr_string.chars().rev().collect::<String>(){
            res = min(res, 1 + helper(s, i, dp))
        }
    }
    dp.insert(ind, res);
    res
}


fn main(){
    assert_eq!(palindrom_partition("ababbbabbababa".to_string()), 3);
    assert_eq!(palindrom_partition("aaabba".to_string()), 1);
    assert_eq!(palindrom_partition("aaa".to_string()), 0);
}