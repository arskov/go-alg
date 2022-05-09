package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode399_0(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}}
	values := []float64{2.0, 2.0, 2.0}
	queries := [][]string{{"a", "b"}, {"b", "a"}, {"b", "c"}, {"c", "b"}, {"c", "d"}, {"d", "c"}, {"a", "d"}, {"d", "a"}, {"x", "x"}, {"x", "y"}}
	expected := []float64{2, 0.5, 2, 0.5, 2, 0.5, 8, 0.125, -1.0, -1.0}
	res := calcEquation_1(equations, values, queries)
	assert.Equal(t, expected, res)
}

func TestLeetCode399_1(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	expected := []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}
	res := calcEquation_1(equations, values, queries)
	assert.Equal(t, expected, res)

}

func TestLeetcode399_3(t *testing.T) {
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}
	values := []float64{1.5, 2.5, 5.0}
	queries := [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}
	expected := []float64{3.75000, 0.40000, 5.00000, 0.20000}
	res := calcEquation_1(equations, values, queries)
	assert.Equal(t, expected, res)
}

func calcEquation_1(equations [][]string, values []float64, queries [][]string) []float64 {
	if equations == nil || values == nil || queries == nil {
		return []float64{}
	}
	graph := make(map[string]map[string]float64)
	for i, v := range equations {
		src := v[0]
		dst := v[1]
		srcDstVal := values[i]
		dstSrcVal := 1.0 / values[i]
		if _, ok := graph[src]; !ok {
			graph[src] = make(map[string]float64)
		}
		if _, ok := graph[dst]; !ok {
			graph[dst] = make(map[string]float64)
		}
		graph[src][dst] = srcDstVal
		graph[dst][src] = dstSrcVal
	}
	res := make([]float64, len(queries))

	var dfs func(string, string, map[string]struct{}) float64

	dfs = func(src, dst string, visited map[string]struct{}) float64 {
		visited[src] = struct{}{}
		adj, ok := graph[src]
		if ok {
			if src == dst {
				return 1.0
			}
			for k, v := range adj {
				if _, ok := visited[k]; !ok {
					res := dfs(k, dst, visited)
					if res > 0 {
						return v * res
					}
				}
			}
		}
		return -1.0
	}

	for i, v := range queries {
		vis := make(map[string]struct{})
		res[i] = dfs(v[0], v[1], vis)
	}

	return res
}
