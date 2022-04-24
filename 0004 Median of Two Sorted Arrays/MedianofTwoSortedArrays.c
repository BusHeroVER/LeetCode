int max(int x, int y) {
    return x > y ? x : y;
}
int min(int x, int y) {
    return x > y ? y : x;
}
double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    if (nums1Size > nums2Size) {
        return findMedianSortedArrays(nums2, nums2Size, nums1, nums1Size);
    }

    int size = nums1Size + nums2Size;
    int mid = (size + 1) >> 1;
    int left = 0;
    int right = nums1Size;
    
    while (left < right) {
        int m1 = (left + right) >> 1;
        int m2 = mid - m1;

        if (nums1[m1] < nums2[m2 -1]) {
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
    int num2 = min(m1 < nums1Size? nums1[m1] : INT_MAX, m2 < nums2Size? nums2[m2] : INT_MAX);
    return (double)(num1 + num2) / 2;
}