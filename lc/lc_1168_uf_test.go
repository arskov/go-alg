package lc

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UnionFind struct {
	root []int
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		root: make([]int, 0),
		rank: make([]int, 0),
	}
	for i := 0; i < n; i++ {
		uf.root = append(uf.root, i)
		uf.rank = append(uf.root, 0)
	}
	return uf
}

func (uf *UnionFind) find(x int) int {
	if x != uf.root[x] {
		uf.root[x] = uf.find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *UnionFind) union(x, y int) bool {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		if uf.rank[rootX] >= uf.rank[rootY] {
			uf.root[rootY] = rootX
			uf.rank[rootX] += uf.rank[rootY]
		} else {
			uf.root[rootX] = rootY
			uf.rank[rootY] += uf.rank[rootX]
		}
		return true
	}
	return false
}

func TestLeetCode1168_1(t *testing.T) {
	wells := []int{1, 2, 2}
	pipes := [][]int{{1, 2, 1}, {2, 3, 1}}
	assert.Equal(t, 3, minCostToSupplyWater_uf(len(wells), wells, pipes))
}

func minCostToSupplyWater_uf(n int, wells []int, pipes [][]int) int {
	orderedEdges := make([][]int, 0)
	orderedEdges = append(orderedEdges, pipes...)
	for i := 0; i < n; i++ {
		orderedEdges = append(orderedEdges, []int{0, i + 1, wells[i]})
	}
	sort.Slice(orderedEdges, func(i, j int) bool {
		return orderedEdges[i][2] < orderedEdges[j][2]
	})
	uf := NewUnionFind(n + 1)
	totalCost := 0
	for _, edge := range orderedEdges {
		if uf.union(edge[0], edge[1]) {
			totalCost += edge[2]
		}
	}
	return totalCost
}
