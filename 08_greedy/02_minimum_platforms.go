import (
	"fmt"
	"slices"
)

func minimumPlatform(arr, dep []int) int {
	slices.Sort(arr)
	slices.Sort(dep)
	n := len(arr)
	i, j, cnt, res := 0, 0, 0, 0
	for i < n && j < n {
		if arr[i] <= dep[j] {
			cnt += 1
			i += 1
		} else {
			cnt -= 1
			j += 1
		}
		res = max(res, cnt)
	}
	return res
}

func main() {
	fmt.Println(minimumPlatform([]int{900, 940, 950, 1100, 1500, 1800}, []int{910, 1200, 1120, 1130, 1900, 2000}) == 3)
	fmt.Println(minimumPlatform([]int{900, 1235, 1100}, []int{1000, 1240, 1200}) == 1)
	fmt.Println(minimumPlatform([]int{1000, 935, 1100}, []int{1200, 1240, 1130}) == 3)
}
