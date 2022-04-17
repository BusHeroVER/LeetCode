#define MAX(a, b) (a > b) ? a : b
int lengthOfLongestSubstring(char * s){
    int s_len = strlen(s);
    if (s_len <= 1)
        return s_len;

    int head = 0, index = 0, ret = 0;
    int map[256] = {0};
    
    while (*s) {
        if (map[*s] != 0) {
            ret = MAX(index - head, ret);
            head = MAX(head, map[*s]);
            
            if (s_len - head < ret)
                return ret;
        }
        map[*s++] = ++index;
    }
    return MAX(s_len - head, ret);
}