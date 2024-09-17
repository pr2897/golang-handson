// package main

// import (
// 	"fmt"
// 	"strings"
// 	"testing"
// )

// func removeProfanity(message *string) {
// 	if message == nil {
// 		return
// 	}

// 	*message = strings.ReplaceAll(*message, "fubb", "****")
// 	*message = strings.ReplaceAll(*message, "shiz", "****")
// 	*message = strings.ReplaceAll(*message, "witch", "*****")
// }

// type car struct {
// 	model string
// 	year  string
// 	cost  float64
// }

// func (c *car) init(model, year string, cost float64) {
// 	c.model = model
// 	c.cost = cost
// 	c.year = year
// }

// func main() {
// 	message := "hello world fubb shiz witch"
// 	removeProfanity(&message)
// 	fmt.Println(message)

// 	removeProfanity(nil)

// 	fmt.Println("tesing")

// 	alto := car{}
// 	alto.init("honda", "2020", 3033.232)
// 	fmt.Println(alto)
// }

// type data struct {
// 	a, b, c, d, e, f, g, h, i, j int64
// }

// var globalPtr *data
// var globalValue data

// func newDataPtr(i int) *data {
// 	data := &data{int64(i), int64(i + 1), int64(i + 2), int64(i + 3), int64(i + 4), int64(i + 5), int64(i + 6), int64(i + 7), int64(i + 8), int64(i + 9)}
// 	return data
// }

// func newData(i int) data {
// 	data := data{int64(i), int64(i + 1), int64(i + 2), int64(i + 3), int64(i + 4), int64(i + 5), int64(i + 6), int64(i + 7), int64(i + 8), int64(i + 9)}
// 	return data
// }

// func BenchmarkProcessValue(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		globalValue = newData(i)
// 	}
// 	// use it to avoid compiler optimization
// 	fmt.Println(globalValue.a)
// }

package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?
func updateBalance(c *customer, t transaction) error {

	if !(t.transactionType == transactionDeposit || t.transactionType == transactionWithdrawal) {
		return errors.New("unknown transaction type")
	}

	if t.transactionType == transactionDeposit {
		c.balance += t.amount
		return nil
	}

	if t.amount > c.balance {
		return errors.New("insufficient funds")
	}

	c.balance -= t.amount
	return nil
}

func main() {
}
