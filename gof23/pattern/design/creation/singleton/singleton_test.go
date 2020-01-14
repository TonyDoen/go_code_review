package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance1(t *testing.T) {
	tf := GetInstance1()
	fmt.Printf("%t\n", tf)
}

func TestGetInstance2(t *testing.T) {
	tf := GetInstance2()
	fmt.Printf("%t\n", tf)
}
