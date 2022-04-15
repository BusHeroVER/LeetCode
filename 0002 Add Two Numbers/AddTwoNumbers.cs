/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null)
            return l2 == null ? null : l2;
        
        ListNode ret = l1;
        int sum = l1.val + l2.val;
        l1.val = sum % 10;
        int tens = sum / 10;

        while (l1.next != null && l2.next != null) {
            l1 = l1.next;
            l2 = l2.next;

            sum = l1.val + l2.val + tens;
            l1.val = sum % 10;
            tens = sum / 10;
        }

        if (l1.next == null) {
            if (l2.next == null) {
                if (tens != 0)
                    l1.next = new ListNode(tens);
                return ret;
            }
            l1.next = l2.next;
        }

        do {
            l1 = l1.next;

            sum = l1.val + tens;
            l1.val = sum % 10;
            tens = sum / 10;

        } while (l1.next != null);
        
        if (tens != 0)
            l1.next = new ListNode(tens);
        return ret;
    }
}
