package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) { // Reading files requires checking most calls for errors. This helper will streamline our error checks below.
	if e != nil {
		panic(e)
	}
}

func main() {
	// read file
	dat, err := ioutil.ReadFile("/tmp/dat") // Perhaps the most basic file reading task is slurping a file’s entire contents into memory.
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat") // You’ll often want more control over how and what parts of a file are read. For these tasks, start by Opening a file to obtain an os.File value.
	check(err)

	b1 := make([]byte, 5) // Read some bytes from the beginning of the file. Allow up to 5 to be read but also note how many actually were read.
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0) // You can also Seek to a known location in the file and Read from there.
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0) // The io package provides some functions that may be helpful for file reading. For example, reads like the ones above can be more robustly implemented with ReadAtLeast.
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0) // There is no built-in rewind, but Seek(0, 0) accomplishes this.
	check(err)

	r4 := bufio.NewReader(f) // The bufio package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close() // Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).

	//
	//
	//
	// write file
	d1 := []byte("hello\ngo\n")
	err = ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err = os.Create("/tmp/dat2") // For more granular writes, open a file for writing.
	check(err)
	defer f.Close() // It’s idiomatic to defer a Close immediately after opening a file.

	d2 := []byte{115, 111, 109, 101, 10} // You can Write byte slices as you’d expect.
	n2, err = f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err = f.WriteString("writes\n") // A WriteString is also available.
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync() // Issue a Sync to flush writes to stable storage.

	w := bufio.NewWriter(f) // bufio provides buffered writers in addition to the buffered readers we saw earlier.
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush() // Use Flush to ensure all buffered operations have been applied to the underlying writer.

}
