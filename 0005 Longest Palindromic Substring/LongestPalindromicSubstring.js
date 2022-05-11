/**
 * @param {string} s
 * @return {string}
 */
//  012...
// ^|a|...|a|$
 var getChar = function(s, index) {
    if (index < 0) return '^';
    if (index > 2 * s.length) return '$';
    if (index % 2 == 0) return '|';
    return s.charAt(parseInt(index / 2));
};
 var longestPalindrome = function(s) {
    var radii = {};
    var maxRad = 0, maxRadC = 0, maxRadRight = 1;

    for (var c = 0; c <= 2 * s.length; c++) {
        if (radii[c] > 0) continue;
        radii[c] = maxRadRight;

        // Palindrome
        for (var l = c - radii[c]; getChar(s, l) == getChar(s, c + (c - l)); l--) {
            radii[c]++;
        }

        // Update radii from center to right
        var maxRight = c + radii[c] - 1;
        for (var l = c - 1; l >= c - radii[c] + 1; l--) {
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

    return s.substr((maxRadC - maxRad + 1) / 2, (maxRad * 2 - 1) / 2);
};