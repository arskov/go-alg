package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode122_0_1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit_2(prices)
	assert.Equal(t, 7, res)
}

func TestLeetcode122_0_2(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	res := maxProfit_2(prices)
	assert.Equal(t, 4, res)
}

func TestLeetcode122_0_3(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	res := maxProfit_2(prices)
	assert.Equal(t, 0, res)
}

func maxProfit_2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	sum := 0
	valleyIdx := 0
	peakIdx := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[peakIdx] {
			peakIdx = i
			continue
		}
		if prices[i] < prices[peakIdx] {
			sum += prices[peakIdx] - prices[valleyIdx]
			valleyIdx = i
			peakIdx = i
		}
	}
	if valleyIdx != peakIdx {
		sum += prices[peakIdx] - prices[valleyIdx]
	}
	return sum
}
