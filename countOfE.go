package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	//var wg sync.WaitGroup
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
	var counter = 0
	for _, fn := range tasks {
		if counter >= E {
			return errors.New("превышено количество ошибок")
		}
		err := fn()
		if err != nil {
			counter++
		}
	}
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
