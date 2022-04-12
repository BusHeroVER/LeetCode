class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        // Brute Force
        // for (int i = 0; i < nums.size() - 1; i++) {
        //     for (int j = i + 1; j < nums.size(); j++) {
        //         if (nums[i] + nums[j] == target)
        //             return {i, j};
        //     }
        // }
        // return {};

        // Two-pass Hash Table
        // unordered_map<int, int> map;
        // for (int i = 0; i < nums.size(); i++) {
        //     map[nums[i]] = i;
        // }

        // for (int i = 0; i < nums.size() - 1; i++) {
        //     int num = target - nums[i];
        //     if (map.count(num) && map[num] != i) {
        //         return {i, map[num]};
        //     }
        // }
        // return {};

        // One-pass Hash Table
        unordered_map<int, int> map;
        for (int i = 0; i < nums.size(); i++) {
            if (map.find(nums[i]) != map.end()) {
                return {map[nums[i]], i};
            }
            map[target - nums[i]] = i;
        }
        return {};
    }
};