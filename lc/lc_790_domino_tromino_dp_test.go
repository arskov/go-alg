package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode790_1_0(t *testing.T) {
	res := numTilings(0)
	assert.Equal(t, 0, res)

	res = numTilings(1)
	assert.Equal(t, 1, res)

	res = numTilings(2)
	assert.Equal(t, 2, res)

	res = numTilings(3)
	assert.Equal(t, 5, res)

	res = numTilings(100)
	assert.Equal(t, 190242381, res)
}

func numTilings(n int) int {
	if n == 0 {
		return 0
	}

	mf := make([]int, n+1)
	mp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		mf[i] = -1
		mp[i] = -1
	}
	var f func(i int) int
	var p func(i int) int

	f = func(i int) int {
		if i < 1 {
			return 0
		}
		if i == 1 || i == 2 {
			return i
		}
		if mf[i] == -1 {
			mf[i] = (f(i-1) + f(i-2) + 2*p(i-1)) % mymath.MOD
		}
		return mf[i]
	}

	p = func(i int) int {
		if i < 2 {
			return 0
		}
		if i == 2 {
			return 1
		}
		if mp[i] == -1 {
			mp[i] = (p(i-1) + f(i-2)) % mymath.MOD
		}
		return mp[i]
	}
	return f(n)
}
