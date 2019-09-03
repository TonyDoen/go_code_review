package easy

import "math"

func MaximizeDistance2ClosestPerson2(seats []int) int {
	n := len(seats)
	var start, res float64
	for i := 0; i < n; i++ {
		if 1 != seats[i] {
			continue
		}
		if 0 == start {
			res = math.Max(res, float64(i) - start)
		} else {
			res = math.Max(res, (float64(i) - start + 1)/2)
		}
		start = float64(i + 1)
	}
	res = math.Max(res, float64(n)-start)
	return int(res)
}