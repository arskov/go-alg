package lc

import (
	"testing"

	"github.com/ArseniKavalchuk/dsa-go/pkg/mymath"
	"github.com/stretchr/testify/assert"
)

func TestLeetcode718_1_0(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	res1 := findLengthBottomUp(nums1, nums2)
	res2 := findLengthTopDown(nums1, nums2)
	assert.Equal(t, 3, res1)
	assert.Equal(t, 3, res2)
}

func TestLeetcode718_1_1(t *testing.T) {
	nums1 := []int{0, 0, 0, 0, 0}
	nums2 := []int{0, 0, 0, 0, 0}
	res1 := findLengthBottomUp(nums1, nums2)
	res2 := findLengthTopDown(nums1, nums2)
	assert.Equal(t, 5, res1)
	assert.Equal(t, 5, res2)
}

func TestLeetcode718_1_3(t *testing.T) {
	nums1 := []int{5, 14, 53, 80, 48}
	nums2 := []int{50, 47, 3, 80, 83}
	res1 := findLengthBottomUp(nums1, nums2)
	res2 := findLengthTopDown(nums1, nums2)
	assert.Equal(t, 1, res1)
	assert.Equal(t, 1, res2)
}

func findLengthTopDown(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	if n == 0 || m == 0 {
		return 0
	}
	l := mymath.Min(n, m)
	memo := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			memo[i][j] = make([]int, l+1)
			for k := 0; k <= l; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	var dp func(i, j, count int) int
	dp = func(i, j, count int) int {
		if i <= 0 || j <= 0 {
			return count
		}
		if memo[i][j][count] == -1 {
			a := count
			if nums1[i-1] == nums2[j-1] {
				a = dp(i-1, j-1, count+1)
			}
			b := dp(i, j-1, 0)
			c := dp(i-1, j, 0)
			memo[i][j][count] = mymath.Max(a, mymath.Max(b, c))
		}
		return memo[i][j][count]
	}
	return dp(n, m, 0)
}

func findLengthBottomUp(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	if n == 0 || m == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	res := 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if nums1[i-1] == nums2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				if dp[i][j] > res {
					res = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return res
}
