package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode740_1_0(t *testing.T) {
	in := []int{3, 4, 2}
	assert.Equal(t, 6, deleteAndEarn(in))
}

func TestLeetcode740_1_1(t *testing.T) {
	in := []int{1, 1, 1, 2, 4, 5, 5, 5, 6}
	assert.Equal(t, 18, deleteAndEarn(in))
}

func TestLeetcode740_1_3(t *testing.T) {
	in := []int{8, 3, 4, 7, 6, 6, 9, 2, 5, 8, 2, 4, 9, 5, 9, 1, 5, 7, 1, 4}
	assert.Equal(t, 61, deleteAndEarn(in))
}

func deleteAndEarn(nums []int) int {
	count := make(map[int]int)
	memo := make(map[int]int)
	maxVal := 0
	for _, v := range nums {
		count[v] += v
		maxVal = mymath.Max(maxVal, v)
	}
	var dfs func(int) int
	dfs = func(num int) int {
		if num == 0 {
			return 0
		}
		if num == 1 {
			return count[num]
		}
		if v, ok := memo[num]; ok {
			return v
		}
		v := mymath.Max(dfs(num-1), count[num]+dfs(num-2))
		memo[num] = v
		return v
	}
	return dfs(maxVal)
}
