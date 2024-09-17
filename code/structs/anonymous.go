package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name != "" && mToSend.sender.number != 0 && mToSend.recipient.name != "" && mToSend.recipient.number != 0 {
		return true
	}

	return false
}

type car struct {
	make    string
	model   string
	doors   int
	mileage int
	// wheel is a field containing an anonymous struct
	wheel struct {
		radius   int
		material string
	}
}

func AnonymousStruct() {
	mycar := struct {
		make string
		year int
	}{
		make: "tesla",
		year: 2023,
	}
	fmt.Println(mycar)
}

/*

Empty Struct
Empty structs are used in Go as a unary value.


// anonymous empty struct type
empty := struct{}{}

// named empty struct type
type emptyStruct struct{}
empty := emptyStruct{}
Copy icon
The cool thing about empty structs is that they're the smallest possible type in Go: they take up zero bytes of memory.


*/
