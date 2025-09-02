import (
	"fmt"
	"sort"
)

func maximumMeetings(start, end []int) int {
	pairs := [][]int{}
	n := len(start)
	for i := range n {
		pairs = append(pairs, []int{start[i], end[i]})
	}
	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	res := 0
	prev_end := -1
	for i := range n {
		if pairs[i][0] > prev_end {
			prev_end = pairs[i][1]
			res += 1
		}
	}
	return res
}

func main() {
	fmt.Println(maximumMeetings([]int{1, 3, 0, 5, 8, 5}, []int{2, 4, 6, 7, 9, 9}) == 4)
	fmt.Println(maximumMeetings([]int{10, 12, 20}, []int{20, 25, 30}) == 1)
	fmt.Println(maximumMeetings([]int{1, 2}, []int{100, 99}) == 1)
}
