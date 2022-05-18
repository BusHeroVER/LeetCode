class Solution {
    public int myAtoi(String s) {
        long ret = 0;
        boolean  isReading = false;
        boolean  isPostive = true;

        for (int i = 0; i < s.length(); i++)
        {
            char c = s.charAt(i);
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

                if (ret < Integer.MIN_VALUE) return Integer.MIN_VALUE;
                else if (Integer.MAX_VALUE < ret) return Integer.MAX_VALUE;
            }
            else break;
        }
        return (int)ret;
    }
}