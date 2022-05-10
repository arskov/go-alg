package mheap

import (
	"container/heap"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	pq := &PQ{}
	j := 0
	for i := 10; i >= 0; i-- {
		*pq = append(*pq, Tuple2{First: i, Second: j})
		j++
	}
	heap.Init(pq)
	for pq.Len() > 0 {
		log.Printf("%+v\n", heap.Pop(pq))
	}
}

func TestPriorityQueuePushPop(t *testing.T) {
	pq := &PQ{}
	j := 1
	for i := 10; i > 0; i-- {
		heap.Push(pq, Tuple2{First: i, Second: j})
		j++
	}
	assert.Equal(t, Tuple2{First: 1, Second: 10}, heap.Pop(pq))
}
