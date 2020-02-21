package main

import (
	"errors"
	"reflect"
)

type Node struct {
	value    interface{}
	previous *Node
	next     *Node
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Set(value interface{}) {
	n.value = value
}

func (n *Node) Previous() *Node {
	return n.previous
}

func (n *Node) Next() *Node {
	return n.next
}

type Queue interface {
	Length() int
	Capacity() int
	Front() *Node
	Rear() *Node
	Enqueue(value interface{}) bool
	Dequeue() interface{}
}

type UniqueQueue struct {
	front    *Node
	rear     *Node
	length   int
	capacity int
	nodeMap  map[interface{}]bool
}

func NewUniqueQueue(capacity int) (*UniqueQueue, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity is less than 0")
	}

	front := &Node{
		value: nil,
	}

	rear := &Node{
		value:    nil,
		previous: front,
	}

	front.next = rear
	nodeMap := make(map[interface{}]bool)

	return &UniqueQueue{
		front:    front,
		rear:     rear,
		capacity: capacity,
		nodeMap:  nodeMap,
	}, nil
}

func (q *UniqueQueue) Enqueue(value interface{}) bool {
	if q.length == q.capacity || value == nil {
		return false
	}

	node := &Node{
		value: value,
	}

	// Ignore uncomparable type.
	if kind := reflect.TypeOf(value).Kind(); kind == reflect.Map || kind == reflect.Slice || kind == reflect.Func {
		return false
	}

	if v, ok := q.nodeMap[value]; ok || v {
		return false
	}

	if q.length == 0 {
		q.front.next = node
	}

	node.previous = q.rear.previous
	node.next = q.rear
	q.rear.previous.next = node
	q.rear.previous = node

	q.nodeMap[value] = true

	q.length++

	return true
}
