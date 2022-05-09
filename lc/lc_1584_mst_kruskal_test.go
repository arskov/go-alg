package lc

import (
	"sort"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/graph"
	"github.com/ArseniKavalchuk/dsa-go/pkg/manhattan"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode1584_1(t *testing.T) {
	in := [][]int{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	}
	res := minCostConnectPoints_1(in)
	assert.Equal(t, 20, res)
}

func minCostConnectPoints_1(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	n := len(points)
	edges := make([]graph.WeightedEdge, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, graph.WeightedEdge{From: i, To: j, Cost: manhattan.Dist(points[i], points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Cost < edges[j].Cost
	})
	root := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = i
	}
	var find func(x int) int
	var union func(x, y int) bool

	find = func(x int) int {
		if x != root[x] {
			root[x] = find(root[x])
		}
		return root[x]
	}

	union = func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			if rank[rootX] >= rank[rootY] {
				root[rootY] = rootX
				rank[rootX] += rank[rootY]
			} else {
				root[rootX] = rootY
				rank[rootY] += rank[rootX]
			}
			return true
		}
		return false
	}

	totalCost := 0
	for _, edge := range edges {
		if union(edge.From, edge.To) {
			totalCost += edge.Cost
		}
	}
	return totalCost
}
