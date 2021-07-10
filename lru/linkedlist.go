package lru

import (
	"github.com/rs/zerolog/log"
)

type Node struct {
	Data string
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Size  int
	Begin *Node
	End   *Node
}

func (ll *DoublyLinkedList) AddNode(data string) {
	log.Debug().Msgf("Adding node: %v", data)
	if ll.Size == 0 {
		newNode := &Node{
			Data: data,
		}
		ll.Begin = newNode
		ll.End = newNode
	} else {
		newNode := &Node{
			Prev: nil,
			Next: ll.Begin,
			Data: data,
		}
		ll.Begin.Prev = newNode
		ll.Begin = newNode
	}
	ll.Size += 1
	log.Debug().Msgf("Current linked list size %v", ll.Size)
}

func (ll *DoublyLinkedList) RemoveNode(data string) {
	log.Debug().Msgf("Removing node: %v", data)
	node := ll.Begin

	// check if data string we're looking for is at the end
	if ll.End != nil {
		if ll.End.Data == data {
			node = ll.End
		}
	}

	for node != nil {
		if node.Data == data {
			if node.Prev != nil {
				if node.Next != nil {
					// target node is between two nodes
					prev := node.Prev
					next := node.Next
					prev.Next = next
					next.Prev = prev
				} else {
					// target node is at the end
					prev := node.Prev
					prev.Next = nil
					ll.End = prev
				}
			} else {
				// target node is at the start
				ll.Begin = node.Next
				if ll.Size == 1 {
					ll.End = ll.Begin
				} else {
					ll.Begin.Prev = nil
				}
			}
			ll.Size -= 1
			log.Debug().Msgf("Current linked list size %v", ll.Size)
			return
		} else {
			node = node.Next
		}
	}
	return
}
