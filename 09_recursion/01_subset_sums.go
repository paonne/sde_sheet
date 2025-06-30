package main

import (
	"fmt"
	"reflect"
	"slices"
)

var res []int

func subsetSums(nums []int) []int {
	helper(0, 0, nums)
	slices.Sort(res)
	return res
}

func helper(ind, curr int, nums []int) {
	if ind == len(nums) {
		res = append(res, curr)
		return
	}
	helper(ind+1, curr, nums)
	helper(ind+1, curr+nums[ind], nums)
}

func main() {
	fmt.Println(reflect.DeepEqual(subsetSums([]int{2, 3}), []int{0, 2, 3, 5}))
	res = []int{}
	fmt.Println(reflect.DeepEqual(subsetSums([]int{1, 2, 1}), []int{0, 1, 1, 2, 2, 3, 3, 4}))
	res = []int{}
	fmt.Println(reflect.DeepEqual(subsetSums([]int{5, 6, 7}), []int{0, 5, 6, 7, 11, 12, 13, 18}))
}
