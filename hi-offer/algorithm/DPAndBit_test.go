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

func TestNumberOf1InInt(t *testing.T) {
	res := NumberOf1InInt(3)
	println(res)
}

func TestPower(t *testing.T) {
	res := Power(5, 2)
	println(res)

	res = Power(5, -2)
	println(res)

	res = Power(0, 0)
	println(res)
}

func TestFind2SingleNumber1(t *testing.T) {
	arr := []int{6, 6, -3, -2, 7, 7, 5, 5, 1, 1, 2, 2}
	res := Find2SingleNumber1(arr)
	for _, v := range res {
		print(v)
		print(" ")
	}
}

func TestFind2SingleNumber2(t *testing.T) {
	arr := []int{6, 6, -3, -2, 7, 7, 5, 5, 1, 1, 2, 2}
	res := Find2SingleNumber2(arr)
	for _, v := range res {
		print(v)
		print(" ")
	}
}
