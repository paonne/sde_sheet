import "fmt"

func minCoins(v int) int {
	coins := []int{1000, 500, 100, 50, 20, 10, 5, 2, 1}
	res := 0
	for _, coin := range coins {
		for v != 0 && coin <= v {
			res += 1
			v -= coin
		}
	}
	return res
}

func main() {
	fmt.Println(minCoins(70) == 2)
	fmt.Println(minCoins(121) == 3)
}
