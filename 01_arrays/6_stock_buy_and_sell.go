func maxProfit(prices []int) int {

	n := len(prices)
	if n == 1 {
		return 0
	}
	mini := prices[0]
	res := 0
	for i := 1; i < n; i++ {
		mini = min(mini, prices[i])
		res = max(res, prices[i]-mini)
	}
	return res
}
