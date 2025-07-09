package main

import (
	"fmt"
	"slices"
)

func maxSumIS(nums []int) int {
	n := len(nums)
	lis := make([]int, n)
	copy(lis, nums)

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				lis[i] = max(lis[i], nums[i]+lis[j])
			}
		}
	}
	return slices.Max(lis)
}

func main() {
	fmt.Println(maxSumIS([]int{1, 101, 2, 3, 100}) == 106)
	fmt.Println(maxSumIS([]int{4, 1, 2, 3}) == 6)
	fmt.Println(maxSumIS([]int{4, 1, 2, 4}) == 7)
}
