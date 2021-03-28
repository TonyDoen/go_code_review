package main

import (
	"errors"
	"fmt"
)

// Argerror
type Argerror struct {
	arg  int
	prob string
}

func (a *Argerror) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.prob)
}

func f24(arg int) (int, error) {
	if 24 == arg {
		return -1, &Argerror{arg, "can't work with 24"}
	}
	return arg + 3, nil
}

//vs.
func f42(arg int) (int, error) {
	if 42 == arg {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

func main() {
	if v, e := f24(24); nil == e {
		fmt.Println(v)
	} else {
		fmt.Println(e)
	}

	if v, e := f42(42); nil == e {
		fmt.Println(v)
	} else {
		fmt.Println(e)
	}

	if v, e := f42(24); nil == e {
		fmt.Println(v)
	} else {
		fmt.Println(e)
	}
}
