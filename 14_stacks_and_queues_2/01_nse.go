import (
	"fmt"
	"reflect"
	"slices"
)

func nextSmallerElement(nums []int) []int {

	stack, res := []int{}, []int{}
	n := len(nums)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res = append(res, -1)
		} else {
			res = append(res, stack[len(stack)-1])
		}
		stack = append(stack, nums[i])
	}
	slices.Reverse(res)
	return res
}

func main() {
	fmt.Println(reflect.DeepEqual(nextSmallerElement([]int{4, 5, 2, 10, 8}), []int{2, 2, -1, 8, -1}))
	fmt.Println(reflect.DeepEqual(nextSmallerElement([]int{3, 2, 1}), []int{2, 1, -1}))
	fmt.Println(reflect.DeepEqual(nextSmallerElement([]int{1, 2, 3}), []int{-1, -1, -1}))
}