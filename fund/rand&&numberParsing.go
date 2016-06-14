package main

import "time"
import "fmt"
import "math/rand"
import "strconv"

func main() { // For example, rand.Intn returns a random int n, 0 <= n < 100.
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())          // rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
	fmt.Print((rand.Float64()*5)+5, ",") // This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano()) // The default number generator is deterministic, so it’ll produce the same sequence of numbers each time by default. To produce varying sequences, give it a seed that changes. Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",") // Call the resulting rand.Rand just like the functions on the rand package.
	fmt.Print(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42) // If you seed a source with the same number, it produces the same sequence of random numbers.
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))

	// number parsing
	fmt.Println()
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("123", 0, 64) // For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
	fmt.Println(i)
	d, _ := strconv.ParseInt("0x1c8", 0, 64) // ParseInt will recognize hex-formatted numbers.
	fmt.Println(d)
	// fmt.Printf("%8x", 456)
	u, _ := strconv.ParseUint("789", 0, 64) // A ParseUint is also available.
	fmt.Println(u)
	k, _ := strconv.Atoi("135") // Atoi is a convenience function for basic base-10 int parsing.
	fmt.Println(k)
	_, e := strconv.Atoi("wat") // Parse functions return an error on bad input.
	fmt.Println(e)
}
