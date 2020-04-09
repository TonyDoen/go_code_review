package medium

import "testing"

func TestCanVisitAllRooms0(t *testing.T) {
	rooms := [][]int{{1}, {2}, {3}, {}}
	res := CanVisitAllRooms0(rooms)
	println(res)

	rooms = [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	res = CanVisitAllRooms0(rooms)
	println(res)
}

func TestCanVisitAllRooms1(t *testing.T) {
	rooms := [][]int{{1}, {2}, {3}, {}}
	res := CanVisitAllRooms1(rooms)
	println(res)

	rooms = [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	res = CanVisitAllRooms1(rooms)
	println(res)
}
