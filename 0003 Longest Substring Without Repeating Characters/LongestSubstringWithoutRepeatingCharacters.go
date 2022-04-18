func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func lengthOfLongestSubstring(s string) int {
	s_len := len(s)
	if s_len <= 1 {
		return s_len
	}

	var (
		head int = 0
		ret  int = 0
	)
	m := map[rune]int{}

	for tail, c := range s {
		if index, ok := m[c]; ok && index != 0 {
			ret = Max(tail-head, ret)
			head = Max(head, index)

			if s_len-head < ret {
				return ret
			}
		}
		m[c] = tail + 1
	}
	return Max(s_len-head, ret)
}