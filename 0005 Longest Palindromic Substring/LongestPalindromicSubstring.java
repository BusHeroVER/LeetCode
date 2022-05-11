class Solution {
    //  012...
    // ^|a|...|a|$
    public Character getChar(String s, int index) {
        if (index < 0) return '^';
        if (index > 2 * s.length()) return '$';
        if (index % 2 == 0) return '|';
        return s.charAt(index / 2);
    }
    public String longestPalindrome(String s) {
        int[] radii = new int[2 * 1000 + 1];
        int maxRad = 0, maxRadC = 0, maxRadRight = 1;

        for (int c = 0; c <= 2 * s.length(); c++) {
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
                radii[c + (c - l)] = Math.min(radii[l], maxRadRight);
            }

            // Update center & rad
            if (radii[c] > maxRad) {
                maxRad = radii[c];
                maxRadC = c;
            }
        }

        return s.substring((maxRadC - maxRad + 1) / 2, (maxRadC - maxRad + 1) / 2 + (maxRad * 2 - 1) / 2);
    }
}