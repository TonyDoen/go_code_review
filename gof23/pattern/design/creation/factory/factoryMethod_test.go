package factory

import (
	"fmt"
	"testing"
)

func TestFactoryMethod1(t *testing.T) {
	bananaPicker := BananaPicker{}

	banana := bananaPicker.MakeFruit("banana")
	fmt.Println(banana)
}
