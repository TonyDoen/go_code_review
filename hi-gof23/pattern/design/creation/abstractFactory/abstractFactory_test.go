package abstractFactory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory1(t *testing.T) {
	factory := PearFactory{}
	pearPicker := factory.MakePicker()
	pear := factory.MakePlant()

	fmt.Println(pear)
	fmt.Println(pearPicker)
}
