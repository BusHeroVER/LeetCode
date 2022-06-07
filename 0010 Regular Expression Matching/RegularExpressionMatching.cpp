class Solution {
public:
    // bool isMatch(string s, string p) {
    //     int m = s.length();
    //     int n = p.length();

    //     if (n == 0) {
    //         return m == 0;
    //     }

    //     //'same char' or '.'
    //     bool matchedCharDot = m != 0 && (s[0] == p[0] || p[0] == '.');
        
    //     //star
    //     bool matchedStar = n >= 2 && p[1] == '*';
    //     if (matchedStar) {
    //         return isMatch(s, p.substr(2, n-2)) || (matchedCharDot && isMatch(s.substr(1, m-1), p));
    //     }

    //     //not star
    //     return matchedCharDot && isMatch(s.substr(1, m-1), p.substr(1, n-1));
    // }
    bool isMatch(string s, string p) {
        int m = s.length();
        int n = p.length();

        bool dp[m+1][n+1];
        
        dp[0][0] = true;
        for (int j = 2; j <= n && p[j-1] == '*'; j += 2) {
            dp[0][j] = true;
        }

        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (p[j-1] != '*') {
                    dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.');
                } else {
                    dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'));
                }
            }
        }
        return dp[m][n];
    }
};