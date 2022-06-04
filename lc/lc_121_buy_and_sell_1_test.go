package lc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode121_0_1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit_1(prices)
	assert.Equal(t, 5, res)
}

func TestLeetcode121_0_2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	res := maxProfit_1(prices)
	assert.Equal(t, 0, res)
}

func maxProfit_1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	minSoFar := math.MaxInt32
	diff := 0
	for _, v := range prices {
		if v < minSoFar {
			minSoFar = v
		}
		if v-minSoFar > diff {
			diff = v - minSoFar
		}
	}
	return diff
}
