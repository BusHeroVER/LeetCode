public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        // Brute Force
        // for (int i = 0; i < nums.length - 1; i++)
        // for (int j = i + 1; j < nums.length; j++)
        //     if (nums[i] + nums[j] == target)
        //         return new int[2] {i, j};
        // return null;

        // Two-pass Hash Table
        // Dictionary<int, int> dic = new Dictionary<int, int>();
        // for (int i = 0; i < nums.Length; i++)
        //     if (!dic.ContainsKey(target - nums[i]))
        //         dic.Add(target - nums[i], i);

        // for (int i = 0; i < nums.Length - 1; i++)
        //     if (dic.ContainsKey(nums[i]))
        //         return new int[2] {i, dic[nums[i]]};
        // return null;

        // One-pass Hash Table
        Dictionary<int, int> dic = new Dictionary<int, int>();
        int key;
        for (int i = 0; i < nums.Length; i++)
            if (dic.ContainsKey(nums[i]))
                return new int[] {dic[nums[i]], i};
            else
            {
                key = target - nums[i];
                if (!dic.ContainsKey(key))
                    dic.Add(key, i);
            }
        return null;
    }
}