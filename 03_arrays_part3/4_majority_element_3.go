func majorityElement(nums []int) []int {
	e1, e2, c1, c2 := 0, 1, 0, 0
	for _, num := range nums {
		if c1 == 0 && e2 != num {
			e1 = num
			c1 = 1
		} else if c2 == 0 && e1 != num {
			e2 = num
			c2 = 1
		} else if num == e1 {
			c1 += 1
		} else if num == e2 {
			c2 += 1
		} else {
			c1 -= 1
			c2 -= 1
		}
	}
	c1, c2 = 0, 0
	for _, num := range nums {
		if num == e1 {
			c1 += 1
		} else if num == e2 {
			c2 += 1
		}
	}
	n := len(nums)
	res := []int{}
	if c1 > n/3 {
		res = append(res, e1)
	}
	if c2 > n/3 {
		res = append(res, e2)
	}
	return res
}
