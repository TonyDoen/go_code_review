package algorithm

import "testing"

func TestFindGreatestSum(t *testing.T) {
	arr := []int{6, -3, -2, 7, -15, 1, 2, 2}
	res := FindGreatestSum(arr)
	println(res)
}

func TestIsMatch(t *testing.T) {
 	s, p := "aa", "a*"
	res := IsMatch(s, p)
	println(res)

	res = IsMatch1(s, p)
	println(res)
}

