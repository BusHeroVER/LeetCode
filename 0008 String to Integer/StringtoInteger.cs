public class Solution {
    public int MyAtoi(string s) {
        long ret = 0;
        bool isReading = false;
        bool isPostive = true;

        for (int i = 0; i < s.Length; i++)
        {
            char c = s[i];
            if (c == ' ')
            {
                if (isReading) break;
            }
            else if (c == '+')
            {
                if (isReading) break;
                else isReading = true;
            }
            else if (c == '-')
            {
                if (isReading) break;
                else
                {
                    isReading = true;
                    isPostive = false;
                }
            }
            else if ('0' <= c && c <= '9')
            {
                isReading = true;

                ret *= 10;
                if (isPostive) ret += c - '0';
                else ret -= c - '0';

                if (ret < int.MinValue) return int.MinValue;
                else if (int.MaxValue < ret) return int.MaxValue;
            }
            else break;
        }
        return (int)ret;
    }
}