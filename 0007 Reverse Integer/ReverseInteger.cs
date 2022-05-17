public class Solution {
    const int max = int.MaxValue / 10;
    const int min = int.MinValue / 10;
    
    public int Reverse(int x) {
        int ret = 0;

        while (x != 0) {
            int digit = x % 10;

            if (ret > max || (ret == max && digit > 7)) {
                return 0;
            }
            if (ret < min || (ret == min && digit < -8)) {
                return 0;
            }

            ret = ret * 10 + digit;
            x /= 10;
        }
        return ret;
    }
}