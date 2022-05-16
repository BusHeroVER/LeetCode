package main

func getChar(s string, index int) rune {
	if index < 0 {
		return '^'
	}
	if index > 2*len(s) {
		return '$'
	}
	if index%2 == 0 {
		return '|'
	}
	return rune(s[index/2])
}
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func longestPalindrome(s string) string {
	var (
		radii       [2*1000 + 1]int
		maxRad      int = 0
		maxRadC     int = 0
		maxRadRight int = 1
	)

	for c := 0; c <= 2*len(s); c++ {
		if radii[c] > 0 {
			continue
		}
		radii[c] = maxRadRight

		// Palindrome
		for l := c - radii[c]; getChar(s, l) == getChar(s, c+(c-l)); l-- {
			radii[c]++
		}

		// Update radii from center to right
		maxRight := c + radii[c] - 1
		for l := c - 1; l >= c-radii[c]+1; l-- {
			maxRadRight = maxRight - (c + (c - l)) + 1
			if radii[l] == maxRadRight {
				break
			}
			radii[c+(c-l)] = Min(radii[l], maxRadRight)
		}

		// Update center & rad
		if radii[c] > maxRad {
			maxRad = radii[c]
			maxRadC = c
		}
	}

	return s[(maxRadC-maxRad+1)/2 : (maxRadC-maxRad+1)/2+(maxRad*2-1)/2]
}
