import (
	"fmt"
	"math"
	"sort"
)

type Item struct {
	val, wt float64
	ratio   float64
}

func fractionalknapsack(val, wt []float64, capacity float64) float64 {
	n := len(wt)
	items := []Item{}
	for i := range n {
		items = append(items, Item{val: val[i], wt: wt[i], ratio: val[i] / wt[i]})
	}
	sort.Slice(items, func(i int, j int) bool {
		return items[i].ratio > items[j].ratio
	})
	res := 0.0
	for _, item := range items {
		if item.wt < capacity {
			res += item.val
			capacity -= item.wt
		} else {
			res += item.val / item.wt * capacity
			break
		}
	}
	return math.Round(res*100) / 100
}

func main() {
	fmt.Println(fractionalknapsack([]float64{60, 100, 120}, []float64{10, 20, 30}, 50) == 240)
	fmt.Println(fractionalknapsack([]float64{500}, []float64{30}, 10) == 166.67)
}