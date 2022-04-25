<?php
class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Float
     */
    function findMedianSortedArrays($nums1, $nums2) {
        $size1 = count($nums1);
        $size2 = count($nums2);
        if ($size1 > $size2) {
            return $this->findMedianSortedArrays($nums2, $nums1);
        }

        $size = $size1 + $size2;
        $mid = ($size + 1) >> 1;
        $left = 0;
        $right = $size1;
        
        while ($left < $right) {
            $m1 = ($left + $right) >> 1;
            $m2 = $mid - $m1;

            if ($nums1[$m1] < $nums2[$m2 -1]) {
                $left = $m1 + 1;
            } else {
                $right = $m1;
            }
        }
        $m1 = $left;
        $m2 = $mid - $m1;

        $num1 = max($m1 > 0? $nums1[$m1 - 1] : PHP_INT_MIN, $m2 > 0? $nums2[$m2 - 1] : PHP_INT_MIN);
        if ($size & 1) {
            return $num1;
        }
        $num2 = min($m1 < $size1? $nums1[$m1] : PHP_INT_MAX, $m2 < $size2? $nums2[$m2] : PHP_INT_MAX);
        return ($num1 + $num2) / 2;
    }
}
?>
