package lc

import (
	"container/heap"
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mheap"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode1631_1_0(t *testing.T) {
	in := [][]int{
		{1, 2, 2},
		{3, 8, 2},
		{5, 3, 5},
	}
	res := minimumEffortPath(in)
	assert.Equal(t, 2, res)
}

func TestLeetcode1631_1_1(t *testing.T) {
	in := [][]int{
		{1, 2, 3},
		{3, 8, 4},
		{5, 3, 5},
	}
	res := minimumEffortPath(in)
	assert.Equal(t, 1, res)
}

func TestLeetcode1631_1_2(t *testing.T) {
	in := [][]int{
		{1, 2, 1, 1, 1},
		{1, 2, 1, 2, 1},
		{1, 2, 1, 2, 1},
		{1, 2, 1, 2, 1},
		{1, 1, 1, 2, 1},
	}
	res := minimumEffortPath(in)
	assert.Equal(t, 0, res)
}

func TestLeetcode1631_1_3(t *testing.T) {
	in := [][]int{
		{1, 10, 6, 7, 9, 10, 4, 9},
	}
	res := minimumEffortPath(in)
	assert.Equal(t, 9, res)
}

func TestLeetcode1631_1_4(t *testing.T) {
	in := [][]int{
		{8, 3, 2, 5, 2, 10, 7, 1, 8, 9},
		{1, 4, 9, 1, 10, 2, 4, 10, 3, 5},
		{4, 10, 10, 3, 6, 1, 3, 9, 8, 8},
		{4, 4, 6, 10, 10, 10, 2, 10, 8, 8},
		{9, 10, 2, 4, 1, 2, 2, 6, 5, 7},
		{2, 9, 2, 6, 1, 4, 7, 6, 10, 9},
		{8, 8, 2, 10, 8, 2, 3, 9, 5, 3},
		{2, 10, 9, 3, 5, 1, 7, 4, 5, 6},
		{2, 3, 9, 2, 5, 10, 2, 7, 1, 8},
		{9, 10, 4, 10, 7, 4, 9, 3, 1, 6},
	}
	res := minimumEffortPath(in)
	assert.Equal(t, 5, res)
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minimumEffortPath(heights [][]int) int {
	rows := len(heights)
	if rows == 0 {
		return 0
	}
	cols := len(heights[0])
	if cols == 0 {
		return 0
	}

	var dir = [][]int{
		{ 0,  1},
		{ 1,  0},
		{ 0, -1},
		{-1,  0},
	}
	dist := make([][]int, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dist[i] = append(dist[i], math.MaxInt32)
		}
	}
	dist[0][0] = 0
	pq := &mheap.PQ{}
	heap.Push(pq, mheap.Tuple3{First: 0, Second: 0, Third: 0})
	for pq.Len() > 0 {
		head := heap.Pop(pq).(mheap.Tuple3)
		x := head.Second
		y := head.Third
		if head.First > dist[x][y] {
			continue
		}
		for _, d := range dir {
			nx := x + d[0]
			ny := y + d[1]
			if nx>=0 && nx<rows && ny>=0 && ny<cols {
				diff := absDiff(heights[x][y], heights[nx][ny])
				maxDiff := maxInt(diff, dist[x][y])
				if maxDiff < dist[nx][ny] {
					dist[nx][ny] = maxDiff
					heap.Push(pq, mheap.Tuple3{First: diff, Second: nx, Third: ny})
				}
			}
		}
	}
	return dist[rows-1][cols-1]
}
