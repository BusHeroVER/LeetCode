class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # Brute Force
        # for i in range(0, len(nums) - 1):
        #     for j in range(i + 1, len(nums)):
        #         if (nums[i] + nums[j] == target):
        #             return [i, j]
        # assert False

        # Two-pass Hash Table
        # dic = dict()
        # for i in range(len(nums)):
        #     dic[target - nums[i]] = i

        # for i in range(0, len(nums) - 1):
        #     if (nums[i] in dic):
        #         return [i, dic[nums[i]]]
        # assert False

        # One-pass Hash Table
        dic = dict()
        for i, n in enumerate(nums):
            if n in dic:
                return dic[n], i
            dic[target - n] = i
        assert False
