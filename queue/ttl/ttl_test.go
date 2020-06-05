// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ttl

import (
	"math/rand"
	"testing"
	"time"

	"github.com/vedranvuk/container/queue/fifo"
)

func TestTTL(t *testing.T) {

	const numitems = 1000

	ids := fifo.New(func(a, b interface{}) bool { return a.(ItemID) == b.(ItemID) })
	onItemTimeout := func(id ItemID, value interface{}) {
		ids.Remove(id)
		// log.Printf("Item '%d' timed out (value: %v)\n", id, value)
	}
	ttl := New(onItemTimeout)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < numitems; i++ {
				id := ttl.Add(i, time.Duration(rand.Intn(1000))*time.Millisecond)
				ids.Push(id)
			}
		}()
	}

	time.Sleep(10 * time.Millisecond)

	toggle := false
	for id, ok := ids.Pop(); ok; id, ok = ids.Pop() {
		toggle = !toggle
		if toggle {
			ttl.Reset(id.(ItemID), time.Duration(rand.Intn(1000))*time.Millisecond)
		} else {
			ttl.Remove(id.(ItemID))
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func BenchmarkTTLAdd(b *testing.B) {
	b.StopTimer()
	ttl := New(nil)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ttl.Add(i, 1*time.Second)
	}
	b.StopTimer()
}

func BenchmarkTTLRemove(b *testing.B) {
	b.StopTimer()
	ttl := New(nil)
	items := make([]ItemID, 0, b.N)
	for i := 0; i < b.N; i++ {
		id := ttl.Add(i, 1*time.Second)
		items = append(items, id)
	}
	b.StartTimer()
	for i := 0; i < len(items); i++ {
		ttl.Remove(items[i])
	}
	b.StopTimer()
}

func BenchmarkTTLRemoveReverse(b *testing.B) {
	b.StopTimer()
	ttl := New(nil)
	items := make([]ItemID, 0, b.N)
	for i := 0; i < b.N; i++ {
		id := ttl.Add(i, 1*time.Second)
		items = append(items, id)
	}
	b.StartTimer()
	for i := len(items) - 1; i != 0; i-- {
		ttl.Remove(items[i])
	}
	b.StopTimer()
}

func BenchmarkTTLReset(b *testing.B) {
	b.StopTimer()
	ttl := New(nil)
	items := make([]ItemID, 0, b.N)
	for i := 0; i < b.N; i++ {
		id := ttl.Add(i, 1*time.Second)
		items = append(items, id)
	}
	b.StartTimer()
	for i := 0; i < len(items); i++ {
		ttl.Reset(items[i], 1*time.Second)
	}
	b.StopTimer()
}

func BenchmarkTTLResetReverse(b *testing.B) {
	b.StopTimer()
	ttl := New(nil)
	items := make([]ItemID, 0, b.N)
	for i := 0; i < b.N; i++ {
		id := ttl.Add(i, 1*time.Second)
		items = append(items, id)
	}
	b.StartTimer()
	for i := len(items) - 1; i != 0; i-- {
		ttl.Reset(items[i], 1*time.Second)
	}
	b.StopTimer()
}
