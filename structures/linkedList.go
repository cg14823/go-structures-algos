package structures

import (
	"errors"
)

type nodeSingle struct {
	Next  *nodeSingle
	Value int
}

// Single Link List. Simple data structure for storing integers
// insert O(N)
// Look up O(N)
// Delete O(N)
type SingleLinkedList struct {
	Head *nodeSingle
	size int
}

// List Iterator used to iterate though the values of a Linked List
type ListIterator struct {
	cursor *nodeSingle
}

func (li *ListIterator) HasNext() bool {
	return li.cursor != nil
}

func (li *ListIterator) Next() int {
	el := li.cursor.Value
	li.cursor = li.cursor.Next
	return el
}

// CreateSingleLinkList returns an empty single linked list
func CreateSingleLinkList() *SingleLinkedList {
	return &SingleLinkedList{
		Head: nil,
		size: 0,
	}
}

// GetSize returns number of nodes in the link list
func (sl *SingleLinkedList) GetSize() int {
	return sl.size
}

// AddLast appends a node to the end of the list
func (sl *SingleLinkedList) AddLast(element int) error {
	node := &nodeSingle{
		Next:  nil,
		Value: element,
	}

	if sl.Head == nil {
		sl.Head = node
		sl.size++
		return nil
	}

	curr := sl.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = node
	sl.size++
	return nil
}

// AddFirst prepends a node to the head of the list
func (sl *SingleLinkedList) AddFirst(element int) error {
	node := &nodeSingle{
		Next:  sl.Head,
		Value: element,
	}
	sl.Head = node
	sl.size++
	return nil
}

// AddAfter appends a node  with the correct key
// @params key - value of node to append after
// @params element - value of new node to be appended
func (sl *SingleLinkedList) AddAfter(key, element int) error {
	curr := sl.Head
	for curr != nil && curr.Value != key {
		curr = curr.Next
	}

	if curr == nil {
		return errors.New("key value not in list")
	}

	node := &nodeSingle{
		Next:  curr.Next,
		Value: element,
	}
	curr.Next = node
	sl.size++
	return nil
}

// AddAfter prepends a node  with the correct key
// @params key - value of node that we will insert in-front of
// @params element - value of new node to be appended
func (sl *SingleLinkedList) AddBefore(key, element int) error {
	if sl.Head == nil {
		return errors.New("key value not in list")
	}

	curr := sl.Head
	prev := sl.Head
	if sl.Head.Value == key {
		node := &nodeSingle{
			Next:  curr,
			Value: element,
		}
		sl.Head = node
		sl.size++
		return nil
	}

	for curr != nil && curr.Value != key {
		prev = curr
		curr = curr.Next
	}

	if curr == nil {
		return errors.New("key value not in list")
	}

	node := &nodeSingle{
		Next:  curr,
		Value: element,
	}

	prev.Next = node
	sl.size++
	return nil
}

// Delete deletes a node. Only delete one node
func (sl *SingleLinkedList) Delete(key int) error {
	// Check head
	if sl.Head == nil {
		return errors.New("can not delete key as it does not exists")
	}

	if sl.Head.Value == key {
		sl.Head = sl.Head.Next
		return nil
	}

	prev := sl.Head
	curr := sl.Head.Next
	for curr != nil && curr.Value != key {
		curr = curr.Next
		prev = curr
	}

	if curr == nil {
		return errors.New("can not delete key as it does not exists")
	}

	prev.Next = curr.Next
	sl.size--
	return nil
}

// Exists will return true if the key is in the link list
func (sl *SingleLinkedList) Exists(key int) bool {
	curr := sl.Head
	for curr != nil && curr.Value != key {
		curr = curr.Next
	}

	return curr != nil
}

// Iterator returns an iterrable object
func (sl *SingleLinkedList) Iterator() *ListIterator {
	return &ListIterator{cursor: sl.Head}
}
