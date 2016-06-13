package main

import (
	"fmt"
	"strings"
)

// In some languages it’s idiomatic to use generic data structures and algorithms. Go does not support generics; in Go it’s common to provide collection functions if and when they are specifically needed for your program and data types.
func Index(vs []string, t string) int { // Returns the first index of the target string t, or -1 if no match is found.
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool { // Returns true if the target string t is in the slice.
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool { // Returns true if one of the strings in the slice satisfies the predicate f.
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool { // Returns true if all of the strings in the slice satisfy the predicate f.
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string { // Returns a new slice containing all strings in the slice that satisfy the predicate f.
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string { // Returns a new slice containing the results of applying the function f to each string in the original slice.
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "peach"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool { return strings.HasPrefix(v, "p") }))
	fmt.Println(All(strs, func(v string) bool { return strings.HasPrefix(v, "p") }))
	fmt.Println(Filter(strs, func(v string) bool { return strings.Contains(v, "e") }))
	fmt.Println(Map(strs, strings.ToUpper))
}

// The above examples all used anonymous functions, but you can also use named functions of the correct type.
