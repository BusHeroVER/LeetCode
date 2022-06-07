package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "aa"))                                  // true
	fmt.Println(isMatch("aa", "a*"))                                  // true
	fmt.Println(isMatch("aa", "a*c*"))                                // true
	fmt.Println(isMatch("ab", ".*"))                                  // true
	fmt.Println(isMatch("cde", ".*e"))                                // true
	fmt.Println(isMatch("abcde", ".*e"))                              // true
	fmt.Println(isMatch("abcde", ".*."))                              // true
	fmt.Println(isMatch("abcde", ".*d.*"))                            // true
	fmt.Println(isMatch("abcde", ".*e.*.*"))                          // true
	fmt.Println(isMatch("aab", "c*a*b"))                              // true
	fmt.Println(isMatch("a", "ab*"))                                  // true
	fmt.Println(isMatch("abbabaaaaaaacaa", "a*.*b.a.*c*a*"))          // true
	fmt.Println(isMatch("abbabaaaaaaacaa", ".*"))                     // true
	fmt.Println(isMatch("abbabaaaaaaacaa", "a*.*b.a.*"))              // true
	fmt.Println(isMatch("abbabaaaaaaacaa", "a*.*b.a.*c*b*a*"))        // true
	fmt.Println(isMatch("abbabaaaaaaacaa", "a*.*b.a.*c*b*a*c*"))      // true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*"))                 // true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*a"))                // false
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab"))               // false
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab."))              // false
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a"))             // false
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*"))           // true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*"))         // true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*a*"))       // true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*"))     // true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*"))   // true
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*b*")) // true

	fmt.Println(isMatch("ddee", "c*a*b"))   // false
	fmt.Println(isMatch("ssippi", "s*p*.")) // false
	fmt.Println(isMatch("abcde", ".*c"))    // false
	fmt.Println(isMatch("aa", "a"))         // false
	fmt.Println(isMatch("abcde", ".*e."))   // false
}

// func isMatch(s string, p string) bool {
// 	lengthS := len(s)
// 	lengthP := len(p)

// 	if lengthP == 0 {
// 		return lengthS == 0
// 	}

// 	//'same char' or '.'
// 	matchedCharDot := lengthS != 0 && (s[0] == p[0] || p[0] == '.')

// 	//star
// 	matchedStar := lengthP >= 2 && p[1] == '*'
// 	if matchedStar {
// 		return (isMatch(s, string(p[2:])) || (matchedCharDot && isMatch(string(s[1:]), p)))
// 	}

// 	//not star
// 	return matchedCharDot && isMatch(string(s[1:]), string(p[1:]))
// }

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	dp := make([][]bool, m+1)

	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for j := 2; j <= n && p[j-1] == '*'; j += 2 {
		dp[0][j] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			} else {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}
	return dp[m][n]
}
