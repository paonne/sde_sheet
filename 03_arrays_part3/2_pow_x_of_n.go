func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := 1.0
	dummy_n := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}(n)
	for dummy_n > 1 {
		if dummy_n%2 == 1 {
			res *= x
			dummy_n -= 1
		} else {
			x *= x
			dummy_n /= 2
		}
	}
	res *= x
	if n < 0 {
		res = 1 / res
	}
	return res
}
