package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(a []int) *ListNode {
	head := &ListNode{Val: a[0]}
	cur := head
	for _, i := range a[1:] {
		newN := &ListNode{Val: i}
		cur.Next = newN
		cur = newN
	}
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	list := createList([]int{1, 2, 3, 3})
	printList(list)
	list = deleteDuplicates(list)
	printList(list)
}
