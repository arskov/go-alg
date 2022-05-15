package lc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode269_1_0(t *testing.T) {
	in := []string{
		"x",
		"y",
		"z",
	}
	assert.Equal(t, "xyz", alienOrder(in))
}

func TestLeetcode269_1_1(t *testing.T) {
	in := []string{
		"x",
		"y",
		"x",
	}
	assert.Equal(t, "", alienOrder(in))
}

func TestLeetcode269_1_2(t *testing.T) {
	in := []string{
		"wrt",
		"wrf",
		"er",
		"ett",
		"rftt",
	}
	assert.Equal(t, "wertf", alienOrder(in))
}

func TestLeetcode269_1_3(t *testing.T) {
	in := []string{
		"ab",
		"adc",
	}
	res := alienOrder(in)
	exp := strings.EqualFold("bacd", res) ||
		strings.EqualFold("abdc", res) ||
		strings.EqualFold("bdca", res) ||
		strings.EqualFold("acbd", res)
	assert.True(t, exp)
}

func alienOrder(words []string) string {
	// check input
	revAdjList := make(map[byte][]byte, 16)
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if _, ok := revAdjList[w[i]]; !ok {
				revAdjList[w[i]] = make([]byte, 0)
			}
		}
	}
	n := len(words)
	for i := 0; i < n-1; i++ {
		word1 := words[i]
		word2 := words[i+1]
		if len(word1) > len(word2) && strings.HasPrefix(word1, word2) {
			return ""
		}
		for j := 0; j < len(word1) && j < len(word2); j++ {
			char1 := word1[j]
			char2 := word2[j]
			if char1 != char2 {
				revAdjList[char2] = append(revAdjList[char2], char1)
				break
			}
		}
	}
	visited := make(map[byte]int)
	order := make([]byte, 0)
	var dfs func(byte) bool
	dfs = func(node byte) bool {
		color, ok := visited[node]
		if !ok {
			visited[node] = 1
			for _, next := range revAdjList[node] {
				res := dfs(next)
				if res {
					return true
				}
			}
			visited[node] = 2
			order = append(order, node)
		} else if color == 1 {
			return true
		}
		return false
	}
	cycle := false
	for k := range revAdjList {
		cycle = cycle || dfs(k)
	}
	if cycle {
		return ""
	} else {
		return string(order)
	}
}
