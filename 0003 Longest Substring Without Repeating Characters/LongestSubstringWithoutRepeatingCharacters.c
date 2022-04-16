#define MAX(a,b) (a>b)?a:b
int lengthOfLongestSubstring(char * s){
    int s_len = strlen(s);
    if (s_len <= 1)
        return s_len;

    int head = 0, ret = 1;
    int map[256] = {0};
    map[s[0]] = 1;
    
    for (int i = 1; i < s_len; i++) {
        char c = s[i];
        if (map[c] != 0) {
            ret = MAX(i - head, ret);
            head = MAX(head, map[c]);;
        }
        map[c] = i + 1;
    }
    return MAX(s_len - head, ret);
}