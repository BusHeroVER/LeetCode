/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		if l2 == nil {
			return nil
		} else {
			return l2
		}
	} else if l2 == nil {
		return l1
	}

	ret := l1
	sum := l1.Val + l2.Val
	l1.Val = sum % 10
	tens := sum / 10

	for l1.Next != nil && l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next

		sum = l1.Val + l2.Val + tens
		l1.Val = sum % 10
		tens = sum / 10
	}

	if l1.Next == nil {
		if l2.Next == nil {
			if tens != 0 {
				l1.Next = new(ListNode)
				l1.Next.Val = tens
				return ret
			}
			return ret
		}
		l1.Next = l2.Next
	}

	for l1.Next != nil {
		l1 = l1.Next

		sum = l1.Val + tens
		l1.Val = sum % 10
		tens = sum / 10
	}

	if tens != 0 {
		l1.Next = new(ListNode)
		l1.Next.Val = tens
	}
	return ret
}