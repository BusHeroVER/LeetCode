public class Solution {
    public int LengthOfLongestSubstring(string s) {
        int s_len = s.Length;
        if (s_len <= 1)
            return s_len;

        int head = 0, ret = 1;
        Dictionary<char, int> dic = new Dictionary<char, int>();
        dic[s[0]] = 1;

        for (int i = 1; i < s_len; i++) {
            char c = s[i];
            if (dic.ContainsKey(c)) {
                ret = Math.Max(i - head, ret);
                head = Math.Max(head, dic[c]);;
            }
            dic[c] = i + 1;
        }
        return Math.Max(s_len - head, ret);
    }
}