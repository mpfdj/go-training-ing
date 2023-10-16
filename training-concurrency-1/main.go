package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	url, close := startServer()
	defer close()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			fmt.Println("index: " + strconv.Itoa(i))
			err := doRequest(url)
			if err != nil {
				panic(err)
			}
		}()

	}

	wg.Wait()
}
