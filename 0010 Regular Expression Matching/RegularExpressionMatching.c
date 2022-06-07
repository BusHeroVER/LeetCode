// bool isMatch(char * s, char * p){
//     int m = strlen(s);
//     int n = strlen(p);

// 	if (n == 0) {
// 		return m == 0;
// 	}

// 	// 'same char' or '.'
//     bool matchedCharDot = m != 0 && (s[0] == p[0] || p[0] == '.');

// 	//star
// 	bool matchedStar = n >= 2 && p[1] == '*';
// 	if (matchedStar) {
// 		char *subs = (char *) calloc(m, sizeof(char));
//     	strncpy(subs, s+1, m);
        
// 		char *subp = (char *) calloc(n-1, sizeof(char));
//     	strncpy(subp, p+2, n-1);
        
// 		return (isMatch(s, subp) || (matchedCharDot && isMatch(subs, p)));
// 	}

// 	//not star
// 	char *subs = (char *) calloc(m, sizeof(char));
// 	strncpy(subs, s+1, m);
    
// 	char *subp = (char *) calloc(n, sizeof(char));
// 	strncpy(subp, p+1, n);
//     return matchedCharDot && isMatch(subs, subp);
// }

bool isMatch(char * s, char * p){
    int m = strlen(s);
    int n = strlen(p);

    bool dp[m+1][n+1];
    memset(dp, 0, (m+1) * (n+1));
    
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