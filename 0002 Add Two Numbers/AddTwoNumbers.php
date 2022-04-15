<?php
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        if (is_null($l1))
            return is_null($l2) ? null : $l2;
        else if (is_null($l2))
            return $l1;
        
        $ret = $l1;
        $sum = $l1->val + $l2->val;
        $l1->val = $sum % 10;
        $tens = intval($sum / 10);

        while (!is_null($l1->next) && !is_null($l2->next)) {
            $l1 = $l1->next;
            $l2 = $l2->next;
    
            $sum = $l1->val + $l2->val + $tens;
            $l1->val = $sum % 10;
            $tens = intval($sum / 10);
        }
    
        if (is_null($l1->next)) {
            if (is_null($l2->next)) {
                if ($tens != 0) {
                    $l1->next = new ListNode($tens);
                    return $ret;
                }
                return $ret;
            }
            $l1->next = $l2->next;
        }

        do {
            $l1 = $l1->next;
    
            $sum = $l1->val + $tens;
            $l1->val = $sum % 10;
            $tens = intval($sum / 10);
    
        } while (!is_null($l1->next));
        
        if ($tens != 0) {
            $l1->next = new ListNode($tens);
        }
        return $ret;
    }
}
?>
