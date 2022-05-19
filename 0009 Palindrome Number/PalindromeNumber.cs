public class Solution {
    public bool IsPalindrome(int x) {
        if (x < 0) return false;

        int org = x;
        long reverse = 0;
        while (x != 0) {
            reverse = reverse*10 + x%10;
            x /= 10;
        }
        return org == reverse;
    }
}