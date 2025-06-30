func partition(s string) [][]string {
	res := [][]string{}
	helper(0, []string{}, s, &res)
	return res
}

func helper(ind int, curr []string, s string, res *[][]string) {
	if ind == len(s) {
		curr_copy := make([]string, len(curr))
		copy(curr_copy, curr)
		*res = append(*res, curr_copy)
		return
	}

	for i := ind; i < len(s); i++ {
		curr_string := s[ind : i+1]
		if is_palindrome(curr_string) {
			curr = append(curr, curr_string)
			helper(i+1, curr, s, res)
			curr = curr[:len(curr)-1]
		}
	}
}

func is_palindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}
	i := 0
	for i <= n/2 {
		if s[i] != s[n-i-1] {
			return false
		}
		i += 1
	}
	return true
}