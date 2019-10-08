package easy

import (
	"fmt"
	"testing"
)

func TestRobotSim1(t *testing.T) {
	commands := []int{4, -1, 4, -2, 4}
	obstacles := [][]int{{2, 4}}
	res1 := RobotSim1(commands, obstacles)
	fmt.Printf("%d\n", res1)
}
