package main

import "fmt"

func missingNumber(nums []int) [2]int {
	n := len(nums)
	Sn := n * (n + 1) / 2
	Sn2 := n * (n + 1) * (2*n + 1) / 6
	c_Sn, c_Sn2 := 0, 0
	for _, num := range nums {
		c_Sn += num
		c_Sn2 += num * num
	}
	eq1 := (c_Sn2 - Sn2) / (c_Sn - Sn)
	eq2 := c_Sn - Sn
	return [2]int{(eq1 + eq2) / 2, (eq1 - eq2) / 2}
}

func main() {
	fmt.Println(missingNumber([]int{3, 1, 2, 5, 3}) == [2]int{3, 4})
	fmt.Println(missingNumber([]int{3, 1, 2, 5, 4, 6, 7, 5}) == [2]int{5, 8})
}
