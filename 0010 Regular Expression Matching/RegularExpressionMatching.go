package main

import (
	"fmt"
)

func main() {
	// DPtable = make(map[string]bool)
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

func isMatch(s string, p string) bool {
	switch {
	case p == "":
		return s == ""
	case s == "":
		if 2 <= len(p) && p[1] == '*' {
			return isMatch(s, p[2:])
		} else if 2 == len(p) && p[1] == '*' {
			return true
		}
		return false
	case 2 < len(p) && p[1] == '*':
		for i, _ := range s {
			var temp bool
			if p[0] == '.' {
				temp = true
			} else {
				temp = isMatch(s[:i], p[:2])
			}
			if temp && isMatch(s[i:], p[2:]) {
				return true
			}
		}
		if isMatch(s, p[:2]) && isMatch("", p[2:]) {
			return true
		}
	case 2 == len(p) && p[1] == '*':
		if p[0] == '.' {
			return true
		}
		for _, c := range s {
			if c != rune(p[0]) {
				return false
			}
		}
		return true
	case p[0] != '.' && s[0] != p[0]:
		return false
	}
	return isMatch(s[1:], p[1:])
}

// Sluggish
// func Hash(K0, K1 string) string {
// 	return K0 + "|" + K1
// }

// var DPtable map[string]bool

// func isMatch(s string, p string) bool {
// 	switch {
// 	case p == "":
// 		return s == ""
// 	case s == "":
// 		if 2 <= len(p) && p[1] == '*' {
// 			key := Hash(s, p[2:])
// 			if _, ok := DPtable[key]; !ok {
// 				DPtable[key] = isMatch(s, p[2:])
// 			}
// 			return DPtable[key]
// 		} else if 2 == len(p) && p[1] == '*' {
// 			return true
// 		}
// 		return false
// 	case 2 < len(p) && p[1] == '*':
// 		for i, _ := range s {
// 			var temp bool
// 			if p[0] == '.' {
// 				temp = true
// 			} else {
// 				key := Hash(s[:i], p[:2])
// 				if _, ok := DPtable[key]; !ok {
// 					DPtable[key] = isMatch(s[:i], p[:2])
// 				}
// 				temp = DPtable[key]
// 			}
// 			if temp {
// 				key := Hash(s[i:], p[2:])
// 				if _, ok := DPtable[key]; !ok {
// 					DPtable[key] = isMatch(s[i:], p[2:])
// 				}
// 				if DPtable[key] {
// 					return true
// 				}
// 			}
// 		}
// 		key := Hash(s, p[:2])
// 		if _, ok := DPtable[key]; !ok {
// 			DPtable[key] = isMatch(s, p[:2])
// 		}
// 		if DPtable[key] {
// 			key := Hash("", p[2:])
// 			if _, ok := DPtable[key]; !ok {
// 				DPtable[key] = isMatch("", p[2:])
// 			}
// 			if DPtable[key] {
// 				return true
// 			}
// 		}
// 	case 2 == len(p) && p[1] == '*':
// 		if p[0] == '.' {
// 			return true
// 		}
// 		for _, c := range s {
// 			if c != rune(p[0]) {
// 				return false
// 			}
// 		}
// 		return true
// 	case p[0] != '.' && s[0] != p[0]:
// 		return false
// 	}
// 	key := Hash(s[1:], p[1:])
// 	if _, ok := DPtable[key]; !ok {
// 		DPtable[key] = isMatch(s[1:], p[1:])
// 	}
// 	return DPtable[key]
// }
