<?php
class Solution {
/**
 * @param Integer[] $nums
 * @param Integer $target
 * @return Integer[]
 */
    function twoSum($nums, $target) {
        // Brute Force
        // for ($i = 0; $i < count($nums) - 1; $i++)
        //     for ($j = $i + 1; $j < count($nums); $j++)
        //         if ($nums[$i] + $nums[$j] == $target)
        //             return array($i, $j);
        // return null;

        // Two-pass Hash Table
        // $map = array();
        // for ($i = 0; $i < count($nums); $i++)
        //     $map[$target - $nums[$i]] = $i;

        // for ($i = 0; $i < count($nums) - 1; $i++)
        //     if (array_key_exists($nums[$i], $map))
        //         return array($i, $map[$nums[$i]]);
        // return null;

        // One-pass Hash Table
        $map = [];
        for ($i = 0; $i < count($nums); $i++) {
            $num = $nums[$i];
            if (isset($map[$num]))
                return [$map[$num], $i];
            $map[$target -$nums[$i]] = $i;
        }
        return null;
    }
}
?>
