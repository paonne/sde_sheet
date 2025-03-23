use std::{
    cmp::{max, min},
    collections::HashMap,
};

fn egg_drop(n: i32, k: i32) -> i32 {
    let mut dp: HashMap<(i32, i32), i32> = HashMap::new();
    return helper(n, k, &mut dp);
}

fn helper(n: i32, k: i32, dp: &mut HashMap<(i32, i32), i32>) -> i32 {
    if dp.contains_key(&(n, k)) {
        return *dp.get(&(n, k)).unwrap();
    }
    if n == 1 || k <= 0 {
        return k;
    }
    let mut res = 100;
    for i in 1..k + 1 {
        let not_break = helper(n, k - i, dp);
        let breaks = helper(n - 1, i - 1, dp);
        res = min(res, 1 + max(breaks, not_break));
    }
    dp.insert((n, k), res);
    return res;
}

fn main() {
    assert_eq!(egg_drop(1, 2), 2);
    assert_eq!(egg_drop(10, 5), 3);
    assert_eq!(egg_drop(2, 10), 4);
}
