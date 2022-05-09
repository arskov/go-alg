package lc

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode332_1(t *testing.T) {
	src := [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}
	res := findItinerary(src)
	assert.Equal(t, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}, res)
}

func TestLeetcode332_2(t *testing.T) {
	src := [][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}
	res := findItinerary(src)
	assert.Equal(t, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}, res)
}

func TestLeetcode332_3(t *testing.T) {
	src := [][]string{
		{"JFK", "KUL"},
		{"JFK", "NRT"},
		{"NRT", "JFK"},
	}
	res := findItinerary(src)
	assert.Equal(t, []string{"JFK", "NRT", "JFK", "KUL"}, res)
}

func TestLeetcode332_4(t *testing.T) {
	src := [][]string{
		{"JFK", "KUL"},
		{"KUL", "JFK"},
		{"JFK", "KUL"},
		{"KUL", "JFK"},
	}
	res := findItinerary(src)
	assert.Equal(t, []string{"JFK", "KUL", "JFK", "KUL", "JFK"}, res)
}

func TestLeetcode332_5(t *testing.T) {
	src := [][]string{
		{"JFK", "AAA"},
		{"AAA", "JFK"},
		{"JFK", "BBB"},
		{"JFK", "CCC"},
		{"CCC", "JFK"},
	}
	res := findItinerary(src)
	assert.Equal(t, []string{"JFK", "AAA", "JFK", "CCC", "JFK", "BBB"}, res)
}

type Node struct {
	key string
	vis bool
}

func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return []string{}
	}
	sort.Slice(tickets, func(i, j int) bool {
		first := strings.Compare(tickets[i][0], tickets[j][0])
		if first == 0 {
			second := strings.Compare(tickets[i][1], tickets[j][1])
			return second <= 0
		}
		return first < 0
	})
	totalTickets := len(tickets)
	graph := make(map[string][]*Node)
	for _, edge := range tickets {
		node := &Node{key: edge[1], vis: false}
		graph[edge[0]] = append(graph[edge[0]], node)
	}
	res := make([][]string, 0)
	var dfs func(path []string) bool
	dfs = func(path []string) bool {
		pathLen := len(path)
		last := path[pathLen-1]
		adj := graph[last]
		if pathLen == totalTickets+1 {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			res = append(res, pathCopy)
			return true
		}
		for _, next := range adj {
			if next.vis {
				continue
			}
			path = append(path, next.key)
			next.vis = true
			ok := dfs(path)
			if ok {
				return ok
			}
			path = path[:len(path)-1]
			next.vis = false
		}
		return false
	}
	dfs([]string{"JFK"})
	return res[0]

}
