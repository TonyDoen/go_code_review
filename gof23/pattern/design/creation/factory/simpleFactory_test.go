package factory

import (
	"fmt"
	"testing"
)

func TestSimpleFactory1(t *testing.T) {
	factory := Factory{}

	apple := factory.MakeApple()
	orange := factory.MakeOrange()

	fmt.Println(apple)
	fmt.Println(orange)
}
