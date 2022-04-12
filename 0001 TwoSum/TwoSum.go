func twoSum(nums []int, target int) []int {
	// Brute Force
	// for i := 0; i < len(nums) - 1; i++ {
	//     for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] + nums[j] == target {
	// 			return []int{i, j}
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
	m := map[int]int{}
	for i, num := range nums {
		key := target - num
		if j, ok := m[key]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}