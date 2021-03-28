package template

import "testing"

func TestDemoStudent(t *testing.T) {
	s := NewStudent("张三", 0, "2000010101010101")
	s.PrepareGotoSchool()

	t1 := NewTeacher("里斯", 0, "1988010101010101")
	t1.PrepareGotoSchool()
}
