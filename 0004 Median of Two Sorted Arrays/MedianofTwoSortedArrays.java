class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int size1 = nums1.length;
        int size2 = nums2.length;
        if (size1 > size2) {
            return findMedianSortedArrays(nums2, nums1);
        }

        int size = size1 + size2;
        int mid = (size + 1) >> 1;
        int left = 0;
        int right = size1;
        int m1, m2;

        while (left < right) {
            m1 = (left + right) >> 1;
            m2 = mid - m1;

            if (nums1[m1] < nums2[m2 - 1]) {
                left = m1 + 1;
            } else {
                right = m1;
            }
        }

        m1 = left;
        m2 = mid - m1;
        
        int num1 = Math.max(m1 > 0? nums1[m1 - 1] : Integer.MIN_VALUE, m2 > 0? nums2[m2 - 1] : Integer.MIN_VALUE);
        if ((size & 1) == 1) {
            return num1;
        }
        int num2 = Math.min(m1 < size1? nums1[m1] : Integer.MAX_VALUE, m2 < size2? nums2[m2] : Integer.MAX_VALUE);
        return (double)(num1 + num2) / 2;
    }
}