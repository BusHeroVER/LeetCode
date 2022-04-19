class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        s_len = len(s)
        if s_len <= 1:
            return s_len

        head = 0
        ret = 0
        dic = dict()

        for index, c in enumerate(s):
            if c in dic:
                ret = max(index - head, ret)
                head = max(dic[c] ,head)

                if s_len - head < ret:
                    return ret
            dic[c] = index + 1
        return max(s_len - head ,ret)
