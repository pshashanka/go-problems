package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	addNode := SinglyLinkedListNode{data, nil}
	if position == 0 {
		addNode.next = llist
		return &addNode
	}

	node := llist
	for i := int32(0); i < position-1; i++ {
		node = node.next
	}

	addNode.next = node.next
	node.next = &addNode

	return llist
}
