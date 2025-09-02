import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i += 1
		}
		j += 1
	}
	return i
}