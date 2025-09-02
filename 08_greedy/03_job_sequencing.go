import (
	"fmt"
	"sort"
)

func JobSequencing(deadline, profit []int) [2]int {
	pairs := [][]int{}
	n := len(deadline)
	for i := range n {
		pairs = append(pairs, []int{deadline[i], profit[i]})
	}
	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})
	cnt, total_profit := 0, 0
	jobs := map[int]bool{}
	for i := range n {
		for day := pairs[i][0]; day > 0; day -= 1 {
			if !jobs[day] {
				jobs[day] = true
				cnt += 1
				total_profit += pairs[i][1]
				break
			}
		}
	}
	return [2]int{cnt, total_profit}
}

func main() {
	fmt.Println(JobSequencing([]int{4, 1, 1, 1}, []int{20, 10, 40, 30}) == [2]int{2, 60})
	fmt.Println(JobSequencing([]int{2, 1, 2, 1, 1}, []int{100, 19, 27, 25, 15}) == [2]int{2, 127})
	fmt.Println(JobSequencing([]int{3, 1, 2, 2}, []int{50, 10, 20, 30}) == [2]int{3, 100})
}