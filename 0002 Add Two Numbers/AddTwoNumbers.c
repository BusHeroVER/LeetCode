/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    if (!l1)
        return l2 ? NULL : l2;
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
                l1->next = (struct ListNode*)malloc(sizeof(struct ListNode));
                l1 = l1->next
                l1->val = tens;
                l1->next = NULL;
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
        l1->next = (struct ListNode*)malloc(sizeof(struct ListNode));
        l1 = l1->next
        l1->val = tens;
        l1->next = NULL;
    }
    return ret;
}