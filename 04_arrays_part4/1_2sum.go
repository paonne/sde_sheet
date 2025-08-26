func twoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for ind, num := range nums {
		req := target - num
		if val, ok := seen[req]; ok {
			return []int{val, ind}
		}
		seen[num] = ind
	}
	return []int{}
}
