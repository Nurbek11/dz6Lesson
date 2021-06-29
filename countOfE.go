package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var fns []func() error
	for i := 0; i < 1000; i++ {
		fns = append(fns, hello)
	}
	err := Execute(fns, 10)
	if err != nil {
		fmt.Println(err)
	}
}

func Execute(tasks []func() error, E int) error {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var counter int
	for _, fn := range tasks {
		wg.Add(1)
		mu.Lock()
		if counter >= E {
			return errors.New("exceeded the number of errors")
		}
		mu.Unlock()
		fn := fn
		go func() {
			var err = fn()
			if err != nil {
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				counter++
			}
		}()
	}
	wg.Wait()
	return nil
}

func hello() error {
	min := -1000000
	max := 1000000
	number := rand.Intn(max-min) + min
	// if number is less than 1000
	if number < 1000 {
		return errors.New("number is less than 1000")
	}
	return nil
}
