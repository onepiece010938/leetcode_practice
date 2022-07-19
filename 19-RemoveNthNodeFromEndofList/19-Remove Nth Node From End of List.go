package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil || head == nil {
		return nil
	}
	p1, p2 := head, head
	i := 0
	for p1 != nil {
		// p1 走到最後一個節點
		if p1.Next == nil {
			// p2沒動的情況 代表要刪開頭
			if i < n {
				head = p2.Next
				break
			}
			// p2也走到最後一個的前一個 代表要刪掉最後一個節點(透過指向nil)
			if p2.Next == p1 {
				p2.Next = nil
				break
			} else {
				//正常情況 刪掉p2下一個節點的值
				//先暫存下下個節點
				tmp := p2.Next.Next
				// p2節點直接連到下下節點
				p2.Next = tmp
				// fmt.Println("P2:", p2.Val)
				break
			}

		}
		if i >= n {
			p2 = p2.Next
		}
		p1 = p1.Next
		i++

	}
	return head
}

/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    slow, fast := dummy, dummy

    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next

    return dummy.Next
}
*/
func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}}}}}
	newhead := removeNthFromEnd(&head, 7)
	fmt.Println("ANS:", newhead)
	fmt.Println("ANS:", newhead)
}
