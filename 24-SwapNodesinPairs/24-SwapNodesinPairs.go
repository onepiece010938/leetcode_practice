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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var last_node *ListNode
	var next_pair *ListNode
	new_head := head.Next
	last_node = head
	for head != nil && head.Next != nil {
		next_pair = head.Next.Next
		head.Next.Next = head
		last_node.Next = head.Next

		head.Next = next_pair
		last_node = head
		head = next_pair

		// 剩一個節點
		if next_pair != nil && next_pair.Next == nil {
			break
		}
	}
	return new_head
}

/*
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{
        Next: head,
    }
    //head=list[i]
    //pre=list[i-1]
    pre := dummy
    for head != nil && head.Next != nil {
        pre.Next = head.Next
        next := head.Next.Next
        head.Next.Next = head
        head.Next = next
        //pre=list[(i+2)-1]
        pre = head
        //head=list[(i+2)]
        head = next
    }
    return dummy.Next
}
*/

// Recursion
/*
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	head.Next = head.Next.Next
	tmp.Next = head

	head.Next = swapPairs(head.Next)

	return tmp
}
*/
func main() {
	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}}}}}
	newhead := swapPairs(&head)
	fmt.Println("ANS:", newhead)
	fmt.Println("ANS:", newhead)
}
