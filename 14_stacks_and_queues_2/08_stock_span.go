type StockSpanner struct {
	prices []int
	pge    []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	this.prices = append(this.prices, price)

	for len(this.pge) > 0 && this.prices[this.pge[len(this.pge)-1]] <= price {
		this.pge = this.pge[:len(this.pge)-1]
	}
	res := len(this.prices)
	if len(this.pge) > 0 {
		res = len(this.prices) - this.pge[len(this.pge)-1] - 1
	}
	this.pge = append(this.pge, len(this.prices)-1)
	return res
}