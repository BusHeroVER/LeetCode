/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        if (!l1)
            return l2 ? nullptr : l2;
        else if (!l2)
            return l1;
        
        struct ListNode *ret = l1;
        int sum = l1->val + l2->val;
        l1->val = sum % 10;
        int tens = sum / 10;

        while (l1->next && l2->next) {
            l1 = l1->next;
            l2 = l2->next;

            sum = l1->val + l2->val + tens;
            l1->val = sum % 10;
            tens = sum / 10;
        }

        if (!l1->next) {
            if (!l2->next) {
                if (tens != 0) {
                    l1->next = new ListNode(tens);
                    return ret;
                }
                return ret;
            }
            l1->next = l2->next;
        }

        do {
            l1 = l1->next;

            sum = l1->val + tens;
            l1->val = sum % 10;
            tens = sum / 10;

        } while (l1->next);
        
        if (tens != 0) {
            l1->next = new ListNode(tens);
        }
        return ret;
    }
};