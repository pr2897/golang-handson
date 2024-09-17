package main

import "fmt"

/*
Go is not an object-oriented language. However, embedded structs provide a kind of data-only inheritance that can be useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, but embedded structs are a way to elevate and share fields between struct definitions

Embedded vs nested
Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
Like nested structs, you assign the promoted fields with the embedded struct in a composite literal.
*/

type sender struct {
	userr
	rateLimit int
}

type userr struct {
	name   string
	number int
}

func main() {
	senders := sender{userr: userr{name: "piyush", number: 23}, rateLimit: 342}
	fmt.Println(senders)
}
