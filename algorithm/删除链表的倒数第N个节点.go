package main

import "fmt"

// 删除链表的倒数第N个节点
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
	if head == nil {
		return nil
	}
	if n == 0 {
		return head
	}
	// curr := head
	for i := 0; i <= n; i++ {
		head = head.Next
	}
	return head

}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

func 删除链表的倒数第N个节点() {
	fmt.Println("test")
	// 数组初始化为链表
	arr := []int{1, 2, 3, 4, 5}
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	printList(head)
}
