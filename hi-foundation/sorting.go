package main

import (
	"fmt"
	"sort"
)

type Bylength []string // In order to sort by a custom function in Go, we need a corresponding type. Here we’ve created a ByLength type that is just an alias for the builtin []string type.

func (s Bylength) Len() int { // We implement sort.Interface - Len, Less, and Swap - on our type so we can use the sort package’s generic Sort function.
	return len(s)
}
func (s Bylength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Bylength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Go’s sort package implements sorting for builtins and user-defined types.
func main() {
	// sorting for builtins
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// Sometimes we’ll want to sort a collection by something other than its natural order.
	// suppose we wanted to sort strings by their length instead of alphabetically.
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(Bylength(fruits)) // we can now implement our custom sort by casting the original fruits slice to ByLength, and then use sort.Sort on that typed slice.
	fmt.Println(fruits)

}
