package lc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode787_3_0(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	res := findCheapestPrice_spfa(4, in, 0, 3, 1)
	assert.Equal(t, 700, res)
}

func TestLeetcode787_3_1(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	res := findCheapestPrice_spfa(3, in, 0, 2, 1)
	assert.Equal(t, 200, res)
}

func TestLeetcode787_3_2(t *testing.T) {
	in := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{0, 2, 500},
	}
	res := findCheapestPrice_spfa(3, in, 0, 2, 0)
	assert.Equal(t, 500, res)
}

func TestLeetcode787_3_3(t *testing.T) {
	in := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 1},
		{3, 7, 1},
		{0, 4, 3},
		{4, 5, 3},
		{5, 7, 3},
		{0, 6, 5},
		{6, 7, 100},
		{7, 8, 1},
	}
	res := findCheapestPrice_spfa(9, in, 0, 8, 3)
	assert.Equal(t, 10, res)
}

func TestLeetcode787_3_4(t *testing.T) {
	in := [][]int{
		{1, 2, 100},
		{3, 1, -150},
		{0, 1, 100},
		{2, 3, 100},
		{0, 3, 200},
		{0, 2, 500},
	}
	res := findCheapestPrice_spfa(4, in, 0, 1, 3)
	assert.Equal(t, 50, res)
}

func TestLeetcode787_3_5(t *testing.T) {
	in := [][]int{
		{0, 1, 1},
		{0, 2, 5},
		{1, 2, 1},
		{2, 3, 1},
	}
	res := findCheapestPrice_spfa(4, in, 0, 3, 1)
	assert.Equal(t, 6, res)
}

func findCheapestPrice_spfa(n int, flights [][]int, src int, dst int, k int) int {
	dist := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		for j := 0; j < n; j++ {
			if j == src {
				dist[i] = append(dist[i], 0)
			} else {
				dist[i] = append(dist[i], math.MaxInt32)
			}
		}
	}
	queue := []int{src}
	res := math.MaxInt32
	stops := 0

	for len(queue) > 0 && stops < k+1 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			head := queue[0]
			queue = queue[1:]
			for _, e := range flights {
				s := e[0]
				d := e[1]
				cost := e[2]
				if stops == k && d != dst {
					continue
				}
				if s == head {
					if dist[stops][s]+cost < dist[stops+1][d] {
						dist[stops+1][d] = dist[stops][s] + cost
						queue = append(queue, d)
						if d == dst {
							if dist[stops+1][d] < res {
								res = dist[stops+1][d]
							}
						}
					}
				}
			}
		}
		stops += 1
	}
	if res == math.MaxInt32 {
		return -1
	} else {
		return res
	}
}
