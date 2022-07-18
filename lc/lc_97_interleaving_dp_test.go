package lc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetcode97_1_0(t *testing.T) {
	s1 := ""
	s2 := ""
	s3 := ""
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_1(t *testing.T) {
	s1 := "a"
	s2 := "a"
	s3 := "aa"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_2(t *testing.T) {
	s1 := "a"
	s2 := "b"
	s3 := "ab"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_3(t *testing.T) {
	s1 := "a"
	s2 := "b"
	s3 := "ba"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_4(t *testing.T) {
	s1 := "ab"
	s2 := "cd"
	s3 := "acbd"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_5(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, false, res)
}

func TestLeetcode97_1_6(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_7(t *testing.T) {
	s1 := "a"
	s2 := ""
	s3 := "a"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_8(t *testing.T) {
	s1 := ""
	s2 := "d"
	s3 := "d"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_9(t *testing.T) {
	s1 := "aabd"
	s2 := "adbc"
	s3 := "aabdabcd"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, false, res)
}

func TestLeetcode97_1_10(t *testing.T) {
	s1 := "aabc"
	s2 := "abad"
	s3 := "aabadabc"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func TestLeetcode97_1_11(t *testing.T) {
	s1 := "aabaac"
	s2 := "aadaaeaaf"
	s3 := "aadaaeaabaafaac"
	res := isInterleave(s1, s2, s3)
	assert.Equal(t, true, res)
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	o := len(s3)
	if n+m != o {
		return false
	}
	if n == 0 {
		return s2 == s3
	}
	if m == 0 {
		return s1 == s3
	}
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}
	var dp func(i, j, k int) bool
	dp = func(i, j, k int) bool {
		if i == n {
			return s2[j:] == s3[k:]
		}
		if j == m {
			return s1[i:] == s3[k:]
		}
		if memo[i][j] == -1 {
			a := s1[i]
			b := s2[j]
			c := s3[k]
			res := false
			if c == a {
				res = res || dp(i+1, j, k+1)
			}
			if c == b {
				res = res || dp(i, j+1, k+1)
			}
			if res {
				memo[i][j] = 1
			} else {
				memo[i][j] = 0
			}
		}
		return memo[i][j] == 1
	}
	return dp(0, 0, 0)
}
