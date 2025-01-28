impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut res = nums[0];
        let mut count = 1;

        for &num in &nums[1..] {
            if num == res {
                count += 1
            } else {
                count -= 1;
                if count == 0 {
                    res = num;
                    count = 1;
                }
            }
        }
        res
    }
}
