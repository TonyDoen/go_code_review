package prototype

import (
	"testing"
)

func TestGetInstance2(t *testing.T) {
	var role Role
	role = &RoleChinese{HeadColor: "black", EyesColor: "black", Tall: 170}

	ThisWriter := role.Clone()
	ThisWriter.SetTall(180)

	role.Show()
	ThisWriter.Show()
}