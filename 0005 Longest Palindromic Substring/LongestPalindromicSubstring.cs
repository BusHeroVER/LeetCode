public class Solution {
    //  012...
    // ^|a|...|a|$
    public char getChar(string s, int index) {
        if (index < 0) return '^';
        if (index > 2 * s.Length) return '$';
        if (index % 2 == 0) return '|';
        return s[index / 2];
    }
    public string LongestPalindrome(string s) {
        int[] radii = new int[2 * 1000 + 1];
        int maxRad = 0, maxRadC = 0, maxRadRight = 1;

        for (int c = 0; c <= 2 * s.Length; c++) {
            if (radii[c] > 0) continue;
            radii[c] = maxRadRight;

            // Palindrome
            for (int l = c - radii[c]; getChar(s, l) == getChar(s, c + (c - l)); l--) {
                radii[c]++;
            }

            // Update radii from center to right
            int maxRight = c + radii[c] - 1;
            for (int l = c - 1; l >= c - radii[c] + 1; l--) {
                maxRadRight = maxRight - (c + (c - l)) + 1;
                if (radii[l] == maxRadRight) break;
                radii[c + (c - l)] = Math.Min(radii[l], maxRadRight);
            }

            // Update center & rad
            if (radii[c] > maxRad) {
                maxRad = radii[c];
                maxRadC = c;
            }
        }

        return s.Substring((maxRadC - maxRad + 1) / 2, (maxRad * 2 - 1) / 2);
    }
}