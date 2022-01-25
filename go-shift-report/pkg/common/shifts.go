package common

import (
	"errors"
	"fmt"
)

type Shift struct {
	Start int64
	End   int64
}

type Node struct {
	Next *Node
	Data *Shift
}

type Shifts struct {
	head *Node
}

func (s *Shifts) AddShift(start, end int64) error {
	if start >= end {
		return errors.New("amount of hours cannot be negative number")
	}

	node := s.Last()
	if node != nil && node.Data.End > start {
		return errors.New("invalid time. must be greater then the previous shift")
	}

	node = &Node{}
	node.Data = &Shift{Start: start, End: end}

	if s.head == nil {
		s.head = node
		return nil
	}

	node.Next = s.head
	s.head = node
	return nil
}

func (s *Shifts) AddTime(ts int64) error {
	node := s.Last()

	// handling the case when the list is entirely empty
	if node == nil {
		node := &Node{Data: &Shift{Start: ts}}
		s.head = node
		return nil
	}

	// handling the case when closing the opened time shift
	if node.Data.End == 0 {
		if ts-node.Data.Start <= 0 {
			return fmt.Errorf("amount of hours cannot be negative or 0")
		}

		node.Data.End = ts
		return nil
	}

	// handling the case when opening a new time shift
	next := node
	if node.Data.End > ts {
		return fmt.Errorf("invalid time. must be greater then a previous shift")
	}

	node = &Node{}
	node.Data = &Shift{Start: ts}

	node.Next = next
	s.head = node
	return nil
}

func (s *Shifts) Last() *Node {
	return s.head
}

func (s *Shifts) Size() int {
	node := s.head

	size := 0
	for node != nil {
		size += 1
		node = node.Next
	}
	return size
}

// !%   @
// nil  1 -> 2 -> 3 -> 4 -> nil
func (s *Shifts) Traverse() {
	var next, prev *Node
	curr := s.head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	s.head = prev
}

func (s *Shifts) Next() *Node {
	node := s.head

	if node == nil {
		return node
	}

	s.head = node.Next
	return node
}
