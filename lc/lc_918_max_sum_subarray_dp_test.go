package lc

import (
	"math"
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode918_1_0(t *testing.T) {
	nums := []int{1, -2, 3, -2}
	res := maxSubarraySumCircular(nums)
	assert.Equal(t, 3, res)
}

func TestLeetcode918_1_1(t *testing.T) {
	nums := []int{5, -3, 5}
	res := maxSubarraySumCircular(nums)
	assert.Equal(t, 10, res)
}

func TestLeetcode918_1_2(t *testing.T) {
	nums := []int{-3, -2, -3}
	res := maxSubarraySumCircular(nums)
	assert.Equal(t, -2, res)
}

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	best := math.MinInt32
	current := 0
	for i := 0; i < n; i++ {
		current = nums[i] + mymath.Max(current, 0)
		best = mymath.Max(best, current)
	}

	rightSum := make([]int, n)
	rightSum[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightSum[i] = nums[i] + rightSum[i+1]
	}

	maxRight := make([]int, n)
	maxRight[n-1] = rightSum[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = mymath.Max(maxRight[i+1], rightSum[i])
	}

	leftSum := 0
	for i := 0; i < n-2; i++ {
		leftSum += nums[i]
		best = mymath.Max(best, leftSum+maxRight[i+2])
	}
	return best
}
