import (
	"fmt"
	"math"
)

func NthRoot(n int, m int) int {

	res := -1

	low, high := 0, m
	var mid int

	for low <= high {
		mid = (low + high) / 2

		curr := int(math.Pow(float64(mid), float64(n)))

		if curr == m {
			return mid
		} else if curr < m {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}

func main(){
	fmt.Println(NthRoot(3, 27) == 3)
	fmt.Println(NthRoot(4, 69) == -1)
}