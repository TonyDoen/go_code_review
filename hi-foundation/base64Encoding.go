package main

// This syntax imports the encoding/base64 package with the b64 name instead of the default base64. It’ll save us some space below.

import b64 "encoding/base64"
import "fmt"

func main() {

	data := "abc123!?$*&()'-=@~" // Here’s the string we’ll encode/decode.

	sEnc := b64.StdEncoding.EncodeToString([]byte(data)) // Go supports both standard and URL-compatible base64. Here’s how to encode using the standard encoder. The encoder requires a []byte so we cast our string to that type.
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc) // Decoding may return an error, which you can check if you don’t already know the input to be well-formed.
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data)) // This encodes/decodes using a URL-compatible base64 format.
	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
