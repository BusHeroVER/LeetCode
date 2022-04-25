class Solution {
public:
    int max(int a, int b) {
        return a > b ? a : b;
    }
    int min(int a, int b) {
        return a > b ? b : a;
    }
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int size1 = nums1.size();
        int size2 = nums2.size();
        if (size1 > size2) {
            return findMedianSortedArrays(nums2, nums1);
        }

        int size = size1 + size2;
        int mid = (size + 1) >> 1;
        int left = 0;
        int right = size1;

        while (left < right) {
            int m1 = (left + right) >> 1;
            int m2 = mid - m1;

            if (nums1[m1] < nums2[m2 - 1]) {
                left = m1 + 1;
            } else {
                right = m1;
            }
        }

        int m1 = left;
        int m2 = mid - m1;

        int num1 = max(m1 > 0? nums1[m1 - 1] : INT_MIN, m2 > 0? nums2[m2 - 1] : INT_MIN);
        if (size & 1) {
            return num1;
        }
        int num2 = min(m1 < size1? nums1[m1] : INT_MAX, m2 < size2? nums2[m2] : INT_MAX);
        return (double)(num1 + num2) / 2;
    }
};