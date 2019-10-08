package easy

import "math"

type Set struct {
	Map map[interface{}]struct{} // struct为结构体类型的变量
}

func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.Map[item] = struct{}{} // 空结构体变量的内存占用大小为0
	}
	return nil
}
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.Map[item]
	return ok
}
func NewSet() *Set {
	return &Set{Map: make(map[interface{}]struct{})} // 获取Set的地址 // 声明map类型的数据结构
}

func RobotSim1(commands []int, obstacles [][]int) int {
	var res, x, y, idx int
	obs := NewSet()
	length := len(obstacles)
	for i := 0; i < length; i++ {
		a := obstacles[i]
		tmp := a[0] + '-' + a[1]
		_ = obs.Add(tmp)
	}
	dirX := []int{0, 1, 0, -1}
	dirY := []int{1, 0, -1, 0}

	length = len(commands)
	for i := 0; i < length; i++ {
		command := commands[i]
		if -1 == command {
			idx = (idx + 1) % 4
		} else if -2 == command {
			idx = (idx - 1 + 4) % 4
		} else {
			tmp := (x + dirX[idx]) + '-' + (y + dirY[idx])
			for command > 0 && !obs.Contains(tmp) {
				x += dirX[idx]
				y += dirY[idx]
				command--
			}
		}
		res = int(math.Max(float64(res), float64(x*x+y*y)))
	}
	return res
}
