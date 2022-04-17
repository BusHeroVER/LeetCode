#define MAX(a, b) (a > b) ? a : b
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int s_len = s.length();
        if (s_len <= 1)
            return s_len;
        
        int head = 0, index = 0, ret = 0;
        int map[256] = {0};
        
        while (index < s_len) {
            char c = s.at(index);
            if (map[c] != 0) {
                ret = MAX(index - head, ret);
                head = MAX(head, map[c]);

                if (s_len - head < ret)
                    return ret;
            }
            map[c] = ++index;
        }
        return MAX(s_len - head, ret);
    }
};