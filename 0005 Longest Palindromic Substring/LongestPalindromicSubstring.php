<?php
class Solution {
    //  012...
    // ^|a|...|a|$
    function getChar($s, $index) {
        if ($index < 0) return '^';
        if ($index > 2 * strlen($s)) return '$';
        if ($index % 2 == 0) return '|';
        return $s[$index / 2];
    }
    /**
     * @param String $s
     * @return String
     */
    function longestPalindrome($s) {
        $radii = array_fill(0, 2 * strlen($s), 0);
        $maxRad = 0;
        $maxRadC = 0;
        $maxRadRight = 1;

        for ($c = 0; $c <= 2 * strlen($s); $c++) {
            if ($radii[$c] > 0) continue;
            $radii[$c] = $maxRadRight;

            // Palindrome
            for ($l = $c - $radii[$c]; $this->getChar($s, $l) == $this->getChar($s, $c + ($c - $l)); $l--) {
                $radii[$c]++;
            }

            // Update radii from center to right
            $maxRight = $c + $radii[$c] - 1;
            for ($l = $c - 1; $l >= $c - $radii[$c] + 1; $l--) {
                $maxRadRight = $maxRight - ($c + ($c - $l)) + 1;
                if ($radii[$l] == $maxRadRight) break;
                $radii[$c + ($c - $l)] = min($radii[$l], $maxRadRight);
            }

            // Update center & rad
            if ($radii[$c] > $maxRad) {
                $maxRad = $radii[$c];
                $maxRadC = $c;
            }
        }

        return substr($s, ($maxRadC - $maxRad + 1) / 2, ($maxRad * 2 - 1) / 2);
    }
}
?>
