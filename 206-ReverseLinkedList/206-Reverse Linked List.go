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

/*
func reverseList(head *ListNode) *ListNode {
	// 異常處理
	if head == nil {
		return head
	}
	//先通通指向第一個
	header := head
	prenode := head
	head = reverse(header, prenode)

	return head
}

func reverse(header *ListNode, prenode *ListNode) (newhead *ListNode) {

	// 我不是最後一個的時候
	// 1.我要指向前一個 2.prenode指標換指到我 3.header指到下一個
	if header.Next != nil {
		// 如果我是第一個 我要指向nil 不是指向自己
		if header == prenode {
			//header指到下一個
			header = header.Next
			// 我要指向nil
			prenode.Next = nil
		} else {
			//暫存我的下一個位置
			tmp := header.Next
			//1.我要指向前一個
			header.Next = prenode
			//2.prenode指標換指到我
			prenode = header
			//3.header指到我的下一個
			header = tmp

		}

		return reverse(header, prenode)

	} else {
		// last element
		// 如果只有我一個
		if header == prenode {
			return header
		}
		// 1.我要指向前一個
		header.Next = prenode
		newhead = header
		return newhead

	}
}
*/

func reverseList(head *ListNode) *ListNode {

	//non or 1 node
	if head == nil || head.Next == nil {
		return head
	}

	//over 2 nodes
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func main() {

	head := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	newhead := reverseList(&head)
	fmt.Println("ANS:", newhead)

}
