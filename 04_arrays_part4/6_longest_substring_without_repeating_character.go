func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	left, res, seen := 0, 0, map[rune]int{}
	for i, c := range s {
		if val, ok := seen[c]; ok && val >= left {
			res = max(res, i-left)
			left = val + 1
		}
		seen[c] = i
	}
	return max(res, n-left)
}
