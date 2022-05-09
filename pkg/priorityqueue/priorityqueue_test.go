package priorityqueue

import (
	"container/heap"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	pq := &PriorityQueue{}
	j := 0
	for i := 10; i >= 0; i-- {
		*pq = append(*pq, Tuple2{Key: i, First: j})
		j++
	}
	heap.Init(pq)
	for pq.Len() > 0 {
		log.Printf("%+v\n", heap.Pop(pq))
	}
}

func TestPriorityQueuePushPop(t *testing.T) {
	pq := &PriorityQueue{}
	j := 1
	for i := 10; i > 0; i-- {
		heap.Push(pq, Tuple2{Key: i, First: j})
		j++
	}
	assert.Equal(t, Tuple2{Key: 1, First: 10}, heap.Pop(pq))
}

