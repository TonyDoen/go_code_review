package main

import "fmt"

func main() {
	// array
	var a [5]int
	b := [5]int{1, 2, 3, 4, 5}

	a[4] = 6
	fmt.Println(a, b)

	var tarr [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			tarr[i][j] = i*j + i + j
			fmt.Println("[", i, ",", j, "] =", tarr[i][j])
		}
	}

	fmt.Println(len(a), len(b), len(tarr), len(tarr[0]))

	//string
	sli := make([]string, 3)
	fmt.Println(sli)

	sli[0] = "a"
	sli[1] = "b"
	sli[2] = "c"
	fmt.Println(sli)

	sli = append(sli, "d")
	sli = append(sli, "e", "f", "g")
	fmt.Println(sli)

	slicopy := make([]string, len(sli))
	copy(slicopy, sli)
	fmt.Println(slicopy)

	//slice
	fmt.Println(sli[2:5])
	fmt.Println(sli[:2])
	fmt.Println(sli[2:])

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	tarr2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		ilen := i + 1
		tarr2[i] = make([]int, ilen)
		for j := 0; j < ilen; j++ {
			tarr2[i][j] = i*j + i + j
			// fmt.Println("[", i, ",", j, "] =", tarr[i][j]) //error
		}
	}
	fmt.Println(tarr2, tarr2[2][0], tarr2[2][1], tarr2[2][2])
}
