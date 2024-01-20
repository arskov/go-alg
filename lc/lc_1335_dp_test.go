package lc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode1335_1_0(t *testing.T) {
	jobDifficulty := []int{1, 1, 1}
	d := 3
	res := minDifficulty(jobDifficulty, d)
	assert.Equal(t, 3, res)
}

func TestLeetcode1335_1_1(t *testing.T) {
	jobDifficulty := []int{1, 2, 3}
	d := 3
	res := minDifficulty(jobDifficulty, d)
	assert.Equal(t, 6, res)
}

func TestLeetcode1335_1_2(t *testing.T) {
	jobDifficulty := []int{1, 1, 1}
	d := 2
	res := minDifficulty(jobDifficulty, d)
	assert.Equal(t, 2, res)
}

func TestLeetcode1335_1_3(t *testing.T) {
	jobDifficulty := []int{1, 1, 1}
	d := 4
	res := minDifficulty(jobDifficulty, d)
	assert.Equal(t, -1, res)
}

func TestLeetcode1335_1_4(t *testing.T) {
	jobDifficulty := []int{6, 5, 4, 3, 2, 1}
	d := 2
	res := minDifficulty(jobDifficulty, d)
	assert.Equal(t, 7, res)
}

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if d > n {
		return -1
	}
	hardest := make([]int, n)
	hardest[n-1] = jobDifficulty[n-1]
	for i := n - 2; i >= 0; i-- {
		hardest[i] = max(jobDifficulty[i], hardest[i+1])
	}
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, d+1)
		for j := 0; j < d+1; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(i, day int) int
	dp = func(i, day int) int {
		if day == d {
			return hardest[i]
		}
		if memo[i][day] == -1 {
			best := math.MaxInt32
			hard := 0
			for j := i; j < n-(d-day); j++ {
				hard = max(hard, jobDifficulty[j])
				best = min(best, hard+dp(j+1, day+1))
			}
			memo[i][day] = best
		}
		return memo[i][day]
	}
	return dp(0, 1)
}
