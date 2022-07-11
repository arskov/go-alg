package lc

import (
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode983_0_1(t *testing.T) {
	days := []int{1, 4, 6, 7, 8, 20}
	cost := []int{2, 7, 15}
	res := mincostTickets(days, cost)
	assert.Equal(t, 11, res)
}

func TestLeetcode983_0_2(t *testing.T) {
	days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	cost := []int{2, 7, 15}
	res := mincostTickets(days, cost)
	assert.Equal(t, 17, res)
}

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	if n == 0 {
		return 0
	}
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	durations := []int{1, 7, 30}
	var dp func(i int) int
	dp = func(i int) int {
		if i >= n {
			return 0
		}
		if memo[i] == -1 {
			res := math.MaxInt32
			j := i
			for k := 0; k < 3; k++ {
				for ; j < n && days[j] < days[i]+durations[k]; j++ {
				}
				res = mymath.Min(res, dp(j)+costs[k])
			}
			memo[i] = res
		}
		return memo[i]
	}
	return dp(0)
}
