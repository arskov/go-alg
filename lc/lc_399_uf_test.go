package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode399_4(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}}
	values := []float64{2.0, 2.0, 2.0}
	queries := [][]string{{"a", "b"}, {"b", "a"}, {"b", "c"}, {"c", "b"}, {"c", "d"}, {"d", "c"}, {"a", "d"}, {"d", "a"}, {"x", "x"}, {"x", "y"}}
	expected := []float64{2, 0.5, 2, 0.5, 2, 0.5, 8, 0.125, -1.0, -1.0}
	res := calcEquation_2(equations, values, queries)
	assert.Equal(t, expected, res)
}

func TestLeetCode399_5(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	expected := []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}
	res := calcEquation_2(equations, values, queries)
	assert.Equal(t, expected, res)

}

func TestLeetcode399_6(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values := []float64{1.5, 2.5, 5.0}
	queries := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	expected := []float64{3.75000, 0.40000, 5.00000, 0.20000}
	res := calcEquation_2(equations, values, queries)
	assert.Equal(t, expected, res)
}

type pair struct {
	gid    string
	weight float64
}

func calcEquation_2(equations [][]string, values []float64, queries [][]string) []float64 {
	gidMap := make(map[string]pair)

	var find func(string) pair
	var union func(string, string, float64)
	var contains func(string) bool

	find = func(x string) pair {
		entry, ok := gidMap[x]
		if !ok {
			entry = pair{gid: x, weight: 1.0}
			gidMap[x] = entry
		}
		if entry.gid != x {
			newEntry := find(entry.gid)
			gidMap[x] = pair{
				gid:    newEntry.gid,
				weight: entry.weight * newEntry.weight,
			}
		}
		return gidMap[x]
	}

	union = func(x, y string, f float64) {
		dividendPair := find(x)
		divisorPair := find(y)
		if dividendPair.gid != divisorPair.gid {
			gidMap[dividendPair.gid] = pair{
				divisorPair.gid,
				divisorPair.weight * f / dividendPair.weight,
			}
		}
	}

	contains = func(s string) bool {
		_, ok := gidMap[s]
		return ok
	}

	for i, equation := range equations {
		union(equation[0], equation[1], values[i])
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		dividendKey := q[0]
		divisorKey := q[1]
		if !contains(dividendKey) || !contains(divisorKey) {
			res[i] = -1.0
		} else {
			dividend := find(dividendKey)
			divisor := find(divisorKey)
			if dividend.gid != divisor.gid {
				res[i] = -1.0
			} else {
				res[i] = dividend.weight / divisor.weight
			}
		}
	}
	return res
}
