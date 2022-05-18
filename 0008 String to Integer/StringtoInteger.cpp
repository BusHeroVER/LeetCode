class Solution {
public:
    int myAtoi(string str) {
        char *s = &str[0];
        long ret = 0;
        bool  isReading = false;
        bool  isPostive = true;

        while (*s) {
            if (*s == ' ') {
                if (isReading) {
                    return ret;
                }
            } else if (*s == '+') {
                if (isReading) {
                    return ret;
                } else {
                    isReading = 1;
                }
            } else if (*s == '-') {
                if (isReading) {
                    return ret;
                } else {
                    isReading = 1;
                    isPostive = 0;
                }
            }  else if ('0' <= *s && *s <= '9') {
                isReading = 1;

                ret *= 10;
                if (isPostive) {
                    ret += *s - '0';
                } else {
                    ret -= *s - '0';
                }

                if (ret < INT_MIN) {
                    return INT_MIN;
                } else if (INT_MAX < ret) {
                    return INT_MAX;
                }
            } else {
                return ret;
            }
            *s++;
        }
        return ret;
    }
};