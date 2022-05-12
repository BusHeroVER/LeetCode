# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        if l1 is None:
            return None if l2 is None else l2
        elif l2 is None:
            return l1

        ret = l1
        sum = l1.val + l2.val
        l1.val = sum % 10
        tens = int(sum / 10)

        while l1.next is not None and l2.next is not None:
            l1 = l1.next
            l2 = l2.next

            sum = l1.val + l2.val + tens
            l1.val = sum % 10
            tens = int(sum / 10)
        
        if l1.next is None:
            if l2.next is None:
                if tens != 0:
                    l1.next = ListNode(tens)
                    return ret
                return ret
            l1.next = l2.next
        
        while l1.next is not None:
            l1 = l1.next
    
            sum = l1.val + tens
            l1.val = sum % 10
            tens = int(sum / 10)
    
        if tens != 0:
            l1.next = ListNode(tens)
        return ret