/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func initNumber(vals []int) *ListNode {
	fmt.Printf("init list node with %v\n", vals)
	root := &ListNode{}
	current := root
	for idx, n := range vals {
		current.Val = n
		if idx == len(vals)-1 {
			current.Next = nil
			break
		}
		current.Next = &ListNode{}
		current = current.Next
	}
	return root
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}

	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	buf := []int{}

	i1 := l1
	i2 := l2
	r := 0
	i1n := l1.Val
	i2n := l2.Val
	for {
		sum := i1n + i2n + r
		l := sum % 10
		r = sum / 10
		fmt.Printf("intermediate sum: [%d %d]\n", l, r)
		buf = append(buf, l)

		if i1 != nil && i1.Next != nil {
			i1 = i1.Next
			i1n = i1.Val
		} else {
			i1 = nil
			i1n = 0
		}

		if i2 != nil && i2.Next != nil {
			i2 = i2.Next
			i2n = i2.Val
		} else {
			i2 = nil
			i2n = 0
		}

		if i1 == nil && i2 == nil {
			if r != 0 {
				// do not forget the reminder here ...
				buf = append(buf, r)
			}
			break
		}
	}

	return initNumber(buf)
}

func main() {
	l1 := initNumber([]int{9, 8})

	l2 := initNumber([]int{1})

	addTwoNumbers(l1, l2)
}
