package lc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode139_1_0(t *testing.T) {
	s := "leetcode"
	wordDict := []string{
		"leet",
		"code",
	}
	res := wordBreak(s, wordDict)
	assert.Equal(t, true, res)
}

func TestLeetcode139_1_1(t *testing.T) {
	s := "applepenapple"
	wordDict := []string{
		"apple",
		"pen",
	}
	res := wordBreak(s, wordDict)
	assert.Equal(t, true, res)
}

func TestLeetcode139_1_2(t *testing.T) {
	s := "catsandog"
	wordDict := []string{
		"cats",
		"dog",
		"sand",
		"and",
		"cat",
	}
	res := wordBreak(s, wordDict)
	assert.Equal(t, false, res)
}

func wordBreak(s string, wordDict []string) bool {
	memo := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		memo[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == len(s) {
			return 1
		}
		if memo[i] == -1 {
			can := 0
			subStr := s[i:]
			for _, w := range wordDict {
				if strings.Index(subStr, w) == 0 {
					can |= dfs(i + len(w))
				}
			}
			memo[i] = can
		}
		return memo[i]
	}
	ans := dfs(0)
	return ans == 1
}
