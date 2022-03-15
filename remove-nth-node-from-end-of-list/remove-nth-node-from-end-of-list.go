/*
 * 19. 删除链表的倒数第 N 个结点
 */

package remove_nth_node_from_end_of_list

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curNode := head
	lens := 0
	for curNode != nil {
		lens += 1
		curNode = curNode.Next
	}
	curNode = head
	fmt.Println(lens)
	if lens == n {
		return head.Next
	}
	for i := 0; i < lens - n - 1; i += 1 {
		curNode = curNode.Next
	}
	curNode.Next = curNode.Next.Next
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	fastNode := head
	slowNode := head
	for i := 0; i < n; i += 1 {
		fastNode = fastNode.Next
	}
	if fastNode == nil {
		return head.Next
	}
	for fastNode.Next != nil {
		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}
	slowNode.Next = slowNode.Next.Next
	return head
}
