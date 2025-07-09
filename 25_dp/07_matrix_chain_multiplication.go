package main

import (
	"fmt"
	"math"
)

func matrixMultiplication(nums []int) int {

	type key struct {
		i int
		j int
	}
	dp := map[key]int{}
	var helper func(i, j int) int

	helper = func(i, j int) int {

		if val, ok := dp[key{i, j}]; ok {
			return val
		}
		if i == j {
			return 0
		}
		mini := math.MaxInt

		for k := i; k < j; k++ {
			mini = min(
				mini,
				helper(i, k)+helper(k+1, j)+nums[i-1]*nums[k]*nums[j],
			)
		}
		dp[key{i, j}] = mini
		return mini
	}
	return helper(1, len(nums)-1)
}

func main() {
	fmt.Println(matrixMultiplication([]int{2, 1, 3, 4}) == 20)
	fmt.Println(matrixMultiplication([]int{1, 2, 3, 4, 3}) == 30)
	fmt.Println(matrixMultiplication([]int{3, 4}) == 0)
}
