package manhattan

import "math"

func Dist(a, b []int) int {
	return int(math.Abs(math.Abs(float64(a[0])-float64(b[0])) + math.Abs(float64(a[1])-float64(b[1]))))
}
