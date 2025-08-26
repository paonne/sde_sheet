func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, zeroes := 0, 0, 0
	for i < m+n && j < n {
		if nums1[i] > nums2[j] {
			nums1 = nums1[:len(nums1)-1]
			nums1 = append(nums1[:i], append([]int{nums2[j]}, nums1[i:]...)...)
			j += 1
			zeroes += 1
		}
		i += 1
	}
	for i := m + zeroes; i < m+n; i++ {
		nums1[i] = nums2[j]
		j += 1
	}
}