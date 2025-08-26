import "fmt"

func solution(nums []int) int {
	sums := map[int]int{0: -1}
	curr, res := 0, 0
	for i, num := range nums {
		curr += num
		if val, ok := sums[curr]; ok {
			res = max(res, i-val)
		} else {
			sums[curr] = i
		}
	}
	return res
}

func main() {
	fmt.Println(solution([]int{9, -3, 3, -1, 6, -5}) == 5)
	fmt.Println(solution([]int{6, -2, 2, -8, 1, 7, 4, -10}) == 8)
	fmt.Println(solution([]int{1, 0, -5}) == 1)
	fmt.Println(solution([]int{1, 3, -5, 6, -2}) == 0)
}
