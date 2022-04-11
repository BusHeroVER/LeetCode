/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSum = function(nums, target) {
    // Brute Force
    // for (var i = 0; i < nums.length - 1; i++) {
    //     for (var j = i + 1; j < nums.length; j++) {
    //         if (nums[i] + nums[j] == target)
    //             return [i, j];
    //     }
    // }
    // return undefined;

    // Two-pass Hash Table
    // var dic = {};
    // for (var i = 0; i < nums.length; i++) {
    //     dic[target - nums[i]] = i;
    // }

    // for (var i = 0; i < nums.length - 1; i++) {
    //     if (dic[nums[i]] != undefined)
    //         return [i, dic[nums[i]]];
    // }
    // return undefined;

    // One-pass Hash Table
    var dic = {};
    for (var i = 0; i < nums.length; i++) {
        var num = nums[i];
        if (dic[num] != undefined) {
            return [dic[num], i];
        } else {
            dic[target - num] = i;
        }
    }
    return undefined;
};