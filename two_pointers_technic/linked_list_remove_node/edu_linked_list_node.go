package main

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}
