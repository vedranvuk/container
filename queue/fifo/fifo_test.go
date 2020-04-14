// Copyright 2020 Vedran Vuk. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package fifo

import (
	"math/rand"
	"testing"
)

func comparefunc(a, b interface{}) bool { return a.(int) == b.(int) }

func TestQueue(t *testing.T) {
	const loops = 1000
	queue := New(comparefunc)
	for i := 0; i < loops; i++ {
		queue.Push(i)
	}
}

func BenchmarkQueuePush(b *testing.B) {
	b.StopTimer()
	queue := New(comparefunc)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
}

func BenchmarkQueuePop(b *testing.B) {
	b.StopTimer()
	queue := New(comparefunc)
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Pop()
	}
}

func BenchmarkQueuePeek(b *testing.B) {
	b.StopTimer()
	queue := New(comparefunc)
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Peek()
	}
}

func BenchmarkQueueRemoveFirst(b *testing.B) {
	b.StopTimer()
	queue := New(comparefunc)
	ids := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		queue.Push(i)
		ids = append(ids, i)
	}
	ididx := 0
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		queue.Remove(ids[ididx])
		b.StopTimer()
		ididx++
	}
}

func BenchmarkQueueRemoveLast(b *testing.B) {
	b.StopTimer()
	queue := New(comparefunc)
	ids := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		queue.Push(i)
		ids = append(ids, i)
	}
	ididx := len(ids) - 1
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		queue.Remove(ids[ididx])
		b.StopTimer()
		ididx--
	}
}

func BenchmarkQueueRemoveRandom(b *testing.B) {
	b.StopTimer()
	queue := New(comparefunc)
	ids := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		queue.Push(i)
		ids = append(ids, i)
	}
	for i := 0; i < b.N; i++ {
		i := rand.Intn(len(ids))
		id := ids[i]
		if i < len(ids)-1 {
			ids = append(ids[:i], ids[i+1:]...)
		} else {
			ids = ids[:i]
		}
		b.StartTimer()
		queue.Remove(id)
		b.StopTimer()
	}
}
