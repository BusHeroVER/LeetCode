class Solution {
    public int[] twoSum(int[] nums, int target) {
        // Brute Force
        // for (int i = 0; i < nums.length - 1; i++) {
        //     for (int j = i + 1; j < nums.length; j++) {
        //         if (nums[i] + nums[j] == target) {
        //             return new int[] {i, j};
        //         }
        //     }
        // }
        // return null;

        // Two-pass Hash Table
        // Map<Integer,Integer> map = new HashMap();
        // for (int i = 0; i < nums.length; i++) {
        //     Integer num = target - nums[i];
        //     if (map.get(num) == null) {
        //         map.put(num, i);
        //     }
        // }

        // for (int i = 0; i < nums.length - 1; i++) {
        //     if (map.get(nums[i]) != null) {
        //         return new int[] {i, map.get(nums[i])};
        //     }
        // }
        // return null;

        // One-pass Hash Table
        Map<Integer,Integer> map = new HashMap();
        for (int i = 0; i < nums.length; i++) {
            int num = nums[i];
            if (map.get(num) != null) {
                return new int[] {map.get(num), i};
            }
            map.put(target - num, i);
        }
        return null;
    }
}