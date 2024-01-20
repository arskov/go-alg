package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/graph"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode1690_1_0(t *testing.T) {
	stones := []int{5, 3, 1, 4, 2}
	res := stoneGameVII(stones)
	assert.Equal(t, 6, res)
}

func TestLeetcode1690_1_1(t *testing.T) {
	stones := []int{7, 90, 5, 1, 100, 10, 10, 2}
	res := stoneGameVII(stones)
	assert.Equal(t, 122, res)
}

func stoneGameVII(stones []int) int {
	// 0 - alice, 1 - bob
	pref := make([]int, len(stones)+1)
	for i := 0; i < len(stones); i++ {
		pref[i+1] = pref[i] + stones[i]
	}
	memo := make(map[graph.Pair]int)
	var game func(p int, s int, e int) int
	game = func(p, s, e int) int {
		if s > e {
			return 0
		}
		k := graph.Pair{First: s, Second: e}
		if v, ok := memo[k]; ok {
			return v
		}
		scoreA := pref[e+1] - pref[s+1]
		scoreB := pref[e] - pref[s]
		if p == 0 {
			a := game(p^1, s+1, e) + scoreA
			b := game(p^1, s, e-1) + scoreB
			memo[k] = max(a, b)
		} else {
			a := game(p^1, s+1, e) - scoreA
			b := game(p^1, s, e-1) - scoreB
			memo[k] = min(a, b)
		}
		return memo[k]
	}
	return game(0, 0, len(stones)-1)
}
