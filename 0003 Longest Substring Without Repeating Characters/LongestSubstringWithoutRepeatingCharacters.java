class Solution {
    public int lengthOfLongestSubstring(String s) {
        int s_len = s.length();
        if (s_len <= 1) {
            return s_len;
        }
        
        int head = 0, index = 0, ret = 0;
        Map<Character, Integer> map = new HashMap<Character, Integer>();

        for (int i = 0; i < s_len;) {
            char c = s.charAt(i);
            if (map.get(c) != null) {
                ret = Math.max(i - head, ret);
                head = Math.max(head, map.get(c));

                if (s_len - head < ret) {
                    return ret;
                }
            }
            map.put(c, ++i);
        }
        return Math.max(s_len - head, ret);
    }
}