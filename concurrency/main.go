package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	wait := waitGroups()

	<-wait
}

func waitGroups() <-chan struct{} {
	ch := make(chan struct{}, 1)

	fmt.Println("help")
	fmt.Println("what am i doing here")

	wg := sync.WaitGroup{}

	for _, file := range []string{"file1.csv", "file2.csv"} {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println(file)
			ch <- struct{}{}
		}()
	}
}

func read(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %w", err)
	}

	ch := make(chan []string)

	go func() {
		cr := csv.NewReader(f)

		for {
			record, err := cr.Read()
			if errors.Is(err, io.EOF) {
				close(ch)

				return
			}

			ch <- record
		}
	}()

	return ch, nil
}
