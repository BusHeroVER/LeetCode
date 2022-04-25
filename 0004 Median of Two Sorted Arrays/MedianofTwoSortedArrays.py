class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        size1 = len(nums1)
        size2 = len(nums2)
        if size1 > size2:
            return self.findMedianSortedArrays(nums2, nums1)
        
        size = size1 + size2
        mid = (size + 1) >> 1
        left = 0
        right = size1

        while left < right:
            m1 = (left + right) >> 1
            m2 = mid - m1

            if nums1[m1] < nums2[m2 -1]:
                left = m1 + 1
            else:
                right = m1
        
        m1 = left
        m2 = mid - m1

        num1 = max(nums1[m1 - 1] if m1 > 0 else float('-inf'), nums2[m2 - 1] if m2 > 0 else float('-inf'))
        if size & 1:
            return num1
    
        num2 = min(nums1[m1] if m1 < size1 else float('inf'), nums2[m2] if m2 < size2 else float('inf'))
        return (num1 + num2) / 2