func isValid(s string) bool {

	stack := []rune{}
	close_map := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, c := range s {
		if _, ok := close_map[c]; !ok {
			stack = append(stack, c)
			continue
		}
		n := len(stack)
		if n == 0 {
			return false
		}
		last := stack[n-1]
		stack = stack[:n-1]
		if last != close_map[c] {
			return false
		}
	}
	return len(stack) == 0
}