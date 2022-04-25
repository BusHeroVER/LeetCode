func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1 := len(nums1)
	size2 := len(nums2)
	if size1 > size2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	var (
		size   int = size1 + size2
		mid    int = (size + 1) >> 1
		left   int = 0
		right  int = size1
		m1     int = 0
		m2     int = 0
		left1  int = 0
		left2  int = 0
		right1 int = 0
		right2 int = 0
	)

	for left < right {
		m1 = (left + right) >> 1
		m2 = mid - m1

		if nums1[m1] < nums2[m2-1] {
			left = m1 + 1
		} else {
			right = m1
		}
	}

	m1 = left
	m2 = mid - m1

	if m1 > 0 {
		left1 = nums1[m1-1]
	} else {
		left1 = math.MinInt
	}
	if m2 > 0 {
		left2 = nums2[m2-1]
	} else {
		left2 = math.MinInt
	}
	num1 := Max(left1, left2)
	if size&1 == 1 {
		return float64(num1)
	}
	if m1 < size1 {
		right1 = nums1[m1]
	} else {
		right1 = math.MaxInt
	}
	if m2 < size2 {
		right2 = nums2[m2]
	} else {
		right2 = math.MaxInt
	}
	num2 := Min(right1, right2)
	return float64(num1+num2) / 2
}