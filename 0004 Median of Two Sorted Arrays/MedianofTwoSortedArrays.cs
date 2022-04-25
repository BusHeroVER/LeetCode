public class Solution {
    public double FindMedianSortedArrays(int[] nums1, int[] nums2) {
        int size1 = nums1.Length;
        int size2 = nums2.Length;
        if (size1 > size2)
            return FindMedianSortedArrays(nums2, nums1);

        int size = size1 + size2;
        int mid = (size + 1) >> 1;
        int left = 0;
        int rigth = size1;
        int m1, m2;
        
        while (left < rigth) {
            m1 = (left + rigth) >> 1;
            m2 = mid - m1;

            if (nums1[m1] < nums2[m2 - 1])
                left = m1 + 1;
            else
                rigth = m1;
        }

        m1 = left;
        m2 = mid - m1;

        int num1 = Math.Max(m1 > 0? nums1[m1 - 1] : int.MinValue, m2 > 0? nums2[m2 - 1] : int.MinValue);
        if ((size & 1) == 1)
            return num1;
        int num2 = Math.Min(m1 < size1? nums1[m1] : int.MaxValue, m2 < size2? nums2[m2] : int.MaxValue);
        return (double)(num1 + num2) / 2;
    }
}