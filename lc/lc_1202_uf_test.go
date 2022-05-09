package lc

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode1202_1(t *testing.T) {
	res := smallestStringWithSwaps(
		"dcab",
		[][]int{{0, 3}, {1, 2}, {0, 2}})
	assert.Equal(t, "abcd", res)
}

func TestLeetCode1202_4(t *testing.T) {
	res := smallestStringWithSwaps(
		"dcab",
		[][]int{{0, 3}, {1, 2}})
	assert.Equal(t, "bacd", res)
}

func TestLeetCode1202_2(t *testing.T) {
	res := smallestStringWithSwaps(
		"udyyek",
		[][]int{{3, 3}, {3, 0}, {5, 1}, {3, 1}, {3, 4}, {3, 5}})
	assert.Equal(t, "deykuy", res)
}

func TestLeetCode1202_3(t *testing.T) {
	res := smallestStringWithSwaps("pwqlmqm",
		[][]int{{5, 3}, {3, 0}, {5, 1}, {1, 1}, {1, 5}, {3, 0}, {0, 2}})
	assert.Equal(t, "lpqqmwm", res)
}

func TestLeetCode1202_5(t *testing.T) {
	res := smallestStringWithSwaps(
		"wyjdlsphvgcalv",
		[][]int{{13, 10}, {1, 3}, {9, 11}, {8, 7}, {7, 7}, {6, 3}, {1, 2}, {6, 1}, {8, 5}, {2, 5}, {5, 8}, {8, 11}, {8, 7}, {10, 12}})
	assert.Equal(t, "wadglhjpsvcylv", res)
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	if len(s) == 0 || pairs == nil || len(pairs) == 0 {
		return s
	}
	n := len(s)
	root := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = i
		rank[i] = i
	}

	var find func(x int) int
	var union func(x, y int) bool

	count := n
	tmp := make([]byte, n)

	find = func(x int) int {
		if x == root[x] {
			return x
		}
		root[x] = find(root[x])
		return root[x]
	}
	union = func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			count--
			if rank[rootX] >= rank[rootY] {
				root[rootY] = rootX
				rank[rootX] += rank[rootY]
			} else {
				root[rootX] = rootY
				rank[rootY] += rank[rootX]
			}
		}
		if count == 1 {
			return true
		} else {
			return false
		}
	}
	for _, pair := range pairs {
		if r := union(pair[0], pair[1]); r {
			break
		}
	}
	comp := make(map[int][]int)
	for i := 0; i < n; i++ {
		r := find(i)
		comp[r] = append(comp[r], i)
	}

	for _, idxs := range comp {
		sort.Slice(idxs, func(i, j int) bool {
			return s[idxs[i]] < s[idxs[j]]
		})
	}
	for i := 0; i < n; i++ {
		r := find(i)
		tmp[i] = s[comp[r][0]]
		comp[r] = comp[r][1:]
	}

	return string(tmp)
}
