func celebrity(M [][]int) int {
	i_know := map[int]int{}
	know_me := map[int]int{}

	n, m := len(M), len(M[0])

	for i := range n {
		for j := range m {
			if M[i][j] == 1 {
				i_know[i] += 1
				know_me[j] += 1
			}
		}
	}

	for i := range n {
		if i_know[i] == 0 && know_me[i] == n-1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(celebrity([][]int{{0, 1, 1, 0}, {0, 0, 0, 0}, {1, 1, 0, 0}, {0, 1, 1, 0}}) == 1)
}