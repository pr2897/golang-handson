fanIn: consolidation of multiple channels into one channel by multiplexing each received value.
```go

```

fanout: separting the resp of one channel to multiple channels.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	parent := make(chan int)

	exit := make(chan bool)

	child1 := make(chan int, 1)
	child2 := make(chan int, 1)
	child3 := make(chan int, 1)

	// producer
	go func() {
		i := 0
		for i < 10 {
			parent <- i
			fmt.Printf("Produced: %d\n", i)
			time.Sleep(time.Millisecond * 100)
			i++
		}

		close(parent)
	}()

	// consumer
	go func() {
		for data := range parent {

			select {
			case child1 <- data:
				fmt.Println("Sent to child1: ", data)
			case child2 <- data:
				fmt.Println("sent to child2: ", data)
			case child3 <- data:
				fmt.Println("sent to child3", data)
			}
		}

		close(child1)
		close(child2)
		close(child3)
		exit <- true
	}()

	// child consumers
	go func() {
		for data := range child1 {
			fmt.Println("child1 received: ", data)
		}
	}()

	go func() {
		for data := range child2 {
			fmt.Println("child2 received: ", data)
		}
	}()

	go func() {
		for data := range child3 {
			fmt.Println("child3 received: ", data)
		}
	}()

	<-exit
}

```