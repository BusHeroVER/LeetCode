class Solution:
    #  012...
    # ^|a|...|a|$
    def getChar(self, s, index):
        if index < 0:
            return "^"
        if index > 2 * len(s):
            return "$"
        if index % 2 == 0:
            return "|"
        return s[int(index / 2)]
    def longestPalindrome(self, s: str) -> str:
        radii = [0] * (2 * len(s) + 1)
        maxRad = 0
        maxRadC = 0
        maxRadRight = 1

        for _, c in enumerate(s):
            if radii[c] > 0: continue
            radii[c] = maxRadRight

            # Palindrome
            for l in range(c - radii[c], -1, -1):
                if self.getChar(s, l) != self.getChar(s, (c + (c - l))):
                    break
                radii[c] += 1

            # Update radii from center to right
            maxRight = c + radii[c] - 1
            for l in range(c - 1,  c - radii[c], -1):
                maxRadRight = maxRight - (c + (c - l)) + 1
                if radii[l] == maxRadRight:
                    break
                radii[c + (c - l)] = min(radii[l], maxRadRight)
            

            # Update center & rad
            if radii[c] > maxRad:
                maxRad = radii[c]
                maxRadC = c

        return s[int((maxRadC - maxRad + 1) / 2): int((maxRadC - maxRad + 1) / 2 + (maxRad * 2 - 1) / 2)]