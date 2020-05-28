package composite

import (
	"testing"
)

func TestDemoComposite(t *testing.T) {
	// 简单组件
	dep1 := NewDepartment("dep1")
	dep2 := NewDepartment("dep2")
	dep3 := NewDepartment("dep3")
	dep4 := NewDepartment("dep4")

	// 复合组件
	cop1 := NewCompany("company1")
	cop1.Add(dep1)
	cop1.Add(dep2)
	cop1.Add(dep3)
	cop1.Add(dep4)

	cop1.Inform("on vacation...")

	cop1.Remove(dep3)
	cop1.Inform("dep3 removed")

	check := cop1.Equals(NewCompany("company1"))
	println(check)
}
