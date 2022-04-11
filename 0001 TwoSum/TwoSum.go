func twoSum(nums []int, target int) []int {
	// Brute Force
	// for i := 0; i < len(nums) - 1; i++ {
	//     for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] + nums[j] == target {
	//             return []int{i, j}
	// 		}
	// 	}
	// }
	// return nil

	// Two-pass Hash Table
	// mmap := make(map[int]int)
	// for i := 0; i < len(nums); i++ {
	//     mmap[target - nums[i]] = i
	// }

	// for i := 0; i < len(nums) - 1; i++ {
	// 	if index, exists := mmap[nums[i]]; exists {
	// 		return []int{i, index}
	// 	}
	// }
	// return nil

	// One-pass Hash Table
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{m[target-num], i}
		}
		m[num] = i
	}
	return nil
}