package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func initList(vals []int) *ListNode {
	head := &ListNode{Val: vals[0], Next: nil}
	rtn := head
	for _, val := range vals[1:] {
		next := &ListNode{Val: val, Next: nil}
		head.Next = next
		head = next
	}

	return rtn
}

func printList(h *ListNode) {
	ptr := h
	for {
		if ptr == nil {
			break
		}

		if ptr.Next != nil {
			fmt.Printf("%d->", ptr.Val)

		} else {
			fmt.Printf("%d\n", ptr.Val)
		}
		ptr = ptr.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr := head
	ptrList := []*ListNode{}
	for {
		if ptr == nil {
			break
		}

		ptrList = append(ptrList, ptr)
		ptr = ptr.Next
	}

	totalLen := len(ptrList)

	if totalLen == n && totalLen == 1 {
		return nil
	}

	// n is always valid, so emit the assertion
	preNodeIdx := totalLen - n - 1
	nextNodeIdx := totalLen - n + 1
	if preNodeIdx == -1 {
		// first node is being removed
		head = ptrList[nextNodeIdx]
	} else {
		node := ptrList[preNodeIdx]
		if n != 1 {
			node.Next = ptrList[nextNodeIdx]
		} else {
			node.Next = nil
		}
	}

	return head
}

func main() {
	testList := initList([]int{1, 2})
	printList(testList)
	printList(removeNthFromEnd(testList, 1))
}
