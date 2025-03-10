package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode 
}

var head = ListNode{Val: 1, Next: &two}

var two = ListNode{Val: 2, Next: &three}

var three = ListNode{Val: 3, Next: &four}

var four = ListNode{Val: 4, Next: &five}

var five = ListNode{Val: 5, Next: nil}


func main() {
	fmt.Println(head)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
	fmt.Println(five)
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        tmp := curr
        curr = curr.Next
        tmp.Next = prev
        prev = tmp
    }

    return prev
}