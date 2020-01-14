package builder

import (
	"fmt"
	"testing"
)

func TestBuilder1(t *testing.T) {
	var builder Builder = &CharacterBuilder{}
	var director *Director = &Director{builder}
	var character *Character = director.Create("loader", "AK47")
	fmt.Println(character.GetName() + "," + character.GetArms())
}
