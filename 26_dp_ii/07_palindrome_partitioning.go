package main

import "fmt"

func palindrom_partitioning(s string) int {
	dp := map[int]int{}
	var helper func(ind int) int

	helper = func(ind int) int {
		if val, ok := dp[ind]; ok {
			return val
		}
		if is_palindrom(s[ind:]) {
			return 0
		}
		res := len(s)
		for i := ind + 1; i < len(s); i++ {
			curr := s[ind:i]
			if is_palindrom(curr) {
				res = min(res, 1+helper(i))
			}
		}
		dp[ind] = res
		return res
	}
	return helper(0)
}

func is_palindrom(s string) bool {
	i, n := 0, len(s)
	for i <= n/2 {
		if s[i] != s[n-i-1] {
			return false
		}
		i += 1
	}
	return true
}

func main() {
	fmt.Println(palindrom_partitioning("ababbbabbababa") == 3)
	fmt.Println(palindrom_partitioning("aaabba") == 1)
	fmt.Println(palindrom_partitioning("aaa") == 0)
}
