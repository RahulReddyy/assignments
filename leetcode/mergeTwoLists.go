package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	val  int
	next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	lkdList1 := list1
	lkdList2 := list2

	slice := []int{}
	linkedData := &ListNode{}
	currentNode := linkedData

	for lkdList1 != nil && lkdList2 != nil {
		slice = append(slice, lkdList1.val, lkdList2.val)
		lkdList1 = lkdList1.next
		lkdList2 = lkdList2.next
	}

	sort.Ints(slice)
	for _, value := range slice {
		currentNode.next = &ListNode{val: value}
		currentNode = currentNode.next
	}

	linkedData = linkedData.next

	return linkedData
}
func main() {
	lkdList1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	lkdList2 := &ListNode{5, &ListNode{1, &ListNode{3, nil}}}
	result := mergeTwoLists(lkdList1, lkdList2)
	fmt.Println(result)
}
