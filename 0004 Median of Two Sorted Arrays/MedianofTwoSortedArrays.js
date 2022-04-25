/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
 var findMedianSortedArrays = function(nums1, nums2) {
    var size1 = nums1.length;
    var size2 = nums2.length;

    if (size1 > size2) {
        return findMedianSortedArrays(nums2, nums1);
    }

    var size = size1 + size2;
    var mid = (size + 1) >> 1;
    var left = 0;
    var right = size1;
    
    while (left < right) {
        var m1 = (left + right) >> 1;
        var m2 = mid - m1;

        if (nums1[m1] < nums2[m2 -1]) {
            left = m1 + 1;
        } else {
            right = m1;
        }
    }
    m1 = left;
    m2 = mid - m1;

    var num1 = Math.max(m1 > 0? nums1[m1 - 1] : Number.MIN_SAFE_INTEGER, m2 > 0? nums2[m2 - 1] : Number.MIN_SAFE_INTEGER);
    if (size & 1) {
        return num1;
    }
    var num2 = Math.min(m1 < size1? nums1[m1] : Number.MAX_SAFE_INTEGER, m2 < size2? nums2[m2] : Number.MAX_SAFE_INTEGER);
    return (num1 + num2) / 2;
};