package main

import "fmt"

/*
114. Flatten Binary Tree to Linked List


Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.


Example 1:


Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [0]
Output: [0]


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-100 <= Node.val <= 100


Follow up: Can you flatten the tree in-place (with O(1) extra space)?
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	dfs(root)
}

/*
	 2
	/ \
   3   4

	 2
	/
   3
    \
	 4

	 2
	/ \
	   3
	    \
	     4
*/
func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftLast := dfs(root.Left)
	rightLast := dfs(root.Right)
	if leftLast != nil {
		leftLast.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	if rightLast != nil {
		return rightLast
	}
	if leftLast != nil {
		return leftLast
	}
	return root
}

func main() {
	//Input: root = [1,2,5,3,4,null,6]
	root := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	flatten(&root)
	fmt.Println(&root)
	fmt.Println(root)

	// bShow(&root)
	// Output: [1,null,2,null,3,null,4,null,5,null,6]
}
