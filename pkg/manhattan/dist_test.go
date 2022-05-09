package manhattan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManhattan(t *testing.T) {
	assert.Equal(t, 4, Dist([]int{0, 0}, []int{2, 2}))
	assert.Equal(t, 4, Dist([]int{0, 0}, []int{-2, -2}))
	assert.Equal(t, 8, Dist([]int{-2, -2}, []int{2, 2}))
}
