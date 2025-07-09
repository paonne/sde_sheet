import "math"

func maxProduct(nums []int) int {

	res := math.MinInt32
	pre, suff := 1, 1
	n := len(nums)

	for i, num := range nums {
		pre = pre * num
		suff = suff * nums[n-i-1]

		res = max(res, pre)
		res = max(res, suff)

		if pre == 0 {
			pre = 1
		}
		if suff == 0 {
			suff = 1
		}
	}
	return res
}
