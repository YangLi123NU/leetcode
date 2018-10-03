package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := res
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		l1Val := 0
		l2Val := 0
		sum := 0
		var node *ListNode
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum = l1Val + l2Val + carry
		if sum > 9 {
			carry = 1
			node = &ListNode{
				Val:  sum - 10,
				Next: nil,
			}
			curr.Next = node
			curr = curr.Next
		} else {
			carry = 0
			node = &ListNode{
				Val:  sum,
				Next: nil,
			}
			curr.Next = node
			curr = curr.Next
		}
	}
	return res.Next
}
