<?php
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLongestSubstring($s) {
        $s_len = strlen($s);
        if ($s_len <= 1) {
            return $s_len;
        }

        $head = 0;
        $ret = 0;
        $map = [];

        for ($index = 0; $index < $s_len;) {
            if (isset($map[$s[$index]])) {
                $ret = max($index - $head, $ret);
                $head = max($map[$s[$index]], $head);

                if ($s_len - $head < $ret) {
                    return $ret;
                }
            }
            $map[$s[$index]] = ++$index;
        }
        return max($s_len - $head, $ret);
    }
}
?>
