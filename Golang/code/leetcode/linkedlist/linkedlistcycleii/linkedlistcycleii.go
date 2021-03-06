/*
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.

链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
*/
package linkedlistcycleii

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

func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := visited[head]; ok {
			return head
		}
		visited[head] = true
		head = head.Next
	}
	return nil
}
