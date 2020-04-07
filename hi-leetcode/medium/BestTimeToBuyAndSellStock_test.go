package medium

import "testing"

func TestGetBestTimeToBuyAndSellStock00(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := GetBestTimeToBuyAndSellStock00(prices)
	println(res)
}

func TestGetBestTimeToBuyAndSellStock01(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := GetBestTimeToBuyAndSellStock01(prices)
	println(res)
}

func TestGetBestTimeToBuyAndSellStock10(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4, 7}
	res := GetBestTimeToBuyAndSellStock10(prices)
	println(res)
}

func TestGetBestTimeToBuyAndSellStock20(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4, 7}
	res := GetBestTimeToBuyAndSellStock20(prices)
	println(res)
}

func TestGetBestTimeToBuyAndSellStock30(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4, 7}
	k := 1
	res := GetBestTimeToBuyAndSellStock30(prices, k)
	println(res)
}