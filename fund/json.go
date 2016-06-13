package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct { // We’ll use these two structs to demonstrate encoding and decoding of custom types below.
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() { // First we’ll look at encoding basic data types to JSON strings. Here are some examples for atomic values.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"} // And here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}} // The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{Page: 1, Fruits: []string{"apple", "peach", "pear"}} // You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of Response2 above to see an example of such tags.
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)    // Now let’s look at decoding JSON data into Go values. Here’s an example for a generic data structure.
	var dat map[string]interface{}                    // We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{} will hold a map of strings to arbitrary data types.
	if err := json.Unmarshal(byt, &dat); err != nil { // Here’s the actual decoding, and a check for associated errors.
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) // In order to use the values in the decoded map, we’ll need to cast them to their appropriate type. For example here we cast the value in num to the expected float64 type.
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) // Accessing nested data requires a series of casts.
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}` // We can also decode JSON into custom data types. This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.
	res := Response2{}
	fmt.Println(res)
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout) // In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out. We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
