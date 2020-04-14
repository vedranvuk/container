// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package fifo implements a first-in-first-out queue.
package fifo

import "sync"

// CompareFunc is a prototype of a func that compares two values for equality.
// It is provided to FIFO via New() and must be able to recognize types of
// values contained in the queue and compare them for equality.
//
// Example:
//
//  queue := New(func(a, b interface{})bool {
//		return a.(int) == b.(int)
//  })
//
//  queue.Push(1)
//  queue.Push(2)
//  queue.Push(3)
//  queue.Remove(2)
type CompareFunc func(a, b interface{}) bool

// Item is an item in the Queue.
type Item struct {
	// next is the Item that follows this Item.
	next *Item
	// Value is the Queue Item Value.
	Value interface{}
}

// FIFO is a first-in-first-out queue. It is safe for concurrent use.
type FIFO struct {
	mu     sync.Mutex
	cf     CompareFunc
	first  *Item
	last   *Item
	length int
}

// New returns a new FIFO with specified comparefunc.
func New(comparefunc CompareFunc) *FIFO {
	return &FIFO{mu: sync.Mutex{}, cf: comparefunc}
}

// Len returns the length of the queue.
func (fifo *FIFO) Len() int {
	fifo.mu.Lock()
	defer fifo.mu.Unlock()
	return fifo.length
}

// Push pushes specified value to the back of the queue.
func (fifo *FIFO) Push(value interface{}) {
	fifo.mu.Lock()
	defer fifo.mu.Unlock()

	li := &Item{nil, value}
	if fifo.last == nil {
		fifo.last = li
	} else {
		fifo.last.next = li
		fifo.last = fifo.last.next
	}
	if fifo.first == nil {
		fifo.first = li
	}
	fifo.length++
}

// Pop removes a Value from the front of the queue and returns it and true if
// queue is not empty, otherwise returns nil and false.
func (fifo *FIFO) Pop() (value interface{}, popped bool) {
	fifo.mu.Lock()
	defer fifo.mu.Unlock()

	if fifo.first == nil {
		return nil, false
	}

	value = fifo.first.Value
	popped = true
	fifo.first = fifo.first.next
	if fifo.first == nil {
		fifo.last = nil
	}
	fifo.length--
	return
}

// Peek returns a Value from the front of the queue and true without removing it
// if queue is not empty, otherwise returns nil and false.
func (fifo *FIFO) Peek() (interface{}, bool) {
	fifo.mu.Lock()
	defer fifo.mu.Unlock()

	if fifo.first == nil {
		return nil, false
	}
	return fifo.first.Value, true
}

// Remove removes a Value from the queue and returns true if queue is not empty,
// false otherwise.
func (fifo *FIFO) Remove(value interface{}) bool {
	fifo.mu.Lock()
	defer fifo.mu.Unlock()

	var last *Item
	curr := fifo.first
	for curr != nil {
		if fifo.cf(curr.Value, value) {
			if last != nil {
				last.next = curr.next
			} else {
				fifo.first = curr.next
			}
			if curr.next == nil {
				fifo.last = nil
			}
			fifo.length--
			return true
		}
		last = curr
		curr = curr.next
	}
	return false
}
