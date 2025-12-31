package main

import "fmt"

// 题目描述：删除链表的倒数第N个节点
// 题目链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 获取链表长度
func getListLength(head *ListNode) int {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	return length

}


// 头指针地址虽然不变，但是整个链表结构发生了改变
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n == 0 {
		return head
	}
	length := getListLength(head)

	// 删除头结点
	if n == length {
		head = head.Next
		return head
	}
	currPoint := head

	// 找到倒数第N个节点的前一个节点
	for i := 1; i < length-n; i++ {
		currPoint = currPoint.Next
	}
	currPoint.Next = currPoint.Next.Next

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
	length := getListLength(head)
	fmt.Printf("链表长度：%d\n", length)
	var resHead *ListNode
	resHead = removeNthFromEnd(head, 1)
	printList(resHead)
	resHead = removeNthFromEnd(head, 2)
	printList(resHead)
	resHead = removeNthFromEnd(head, 3)
	printList(resHead)
	resHead = removeNthFromEnd(head, 4)
	printList(resHead)
	resHead = removeNthFromEnd(head, 5)
	printList(resHead)

}
