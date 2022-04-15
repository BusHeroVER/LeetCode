/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
 var addTwoNumbers = function(l1, l2) {
    if (l1 == undefined)
        return l2 == undefined ? undefined : l2;
    else if (l2 == undefined)
        return l1;
    
    var ret = l1;
    var sum = l1.val + l2.val;
    l1.val = sum % 10;
    var tens = Math.trunc(sum / 10);
    
    while (l1.next != null && l2.next != null) {
        l1 = l1.next;
        l2 = l2.next;

        sum = l1.val + l2.val + tens;
        l1.val = sum % 10;
        tens = Math.trunc(sum / 10);
    }
    
    if (l1.next == null) {
        if (l2.next == null) {
            if (tens != 0) {
                l1.next = new ListNode(tens);
                return ret;
            }
            return ret;
        }
        l1.next = l2.next;
    }

    do {
        l1 = l1.next;
        
        sum = l1.val + tens;
        l1.val = sum % 10;
        tens = Math.trunc(sum / 10);
        
    } while (l1.next != null);
    
    if (tens != 0) {
        l1.next = new ListNode(tens);
    }
    return ret;
};