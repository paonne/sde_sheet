import "slices"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	lis := make([]int, n)
	for i := range n {
		lis[i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				lis[i] = max(lis[i], 1+lis[j])
			}
		}
	}
	return slices.Max(lis)
}