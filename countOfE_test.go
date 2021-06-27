package main

import (
	"fmt"
	"testing"
)

func CountOfE(t *testing.T) {
	var fns []func() error
	var fns2 []func() error
	for i := 0; i < 1000; i++ {
		fns = append(fns, hello)
	}
	for i := 0; i < 10; i++ {
		fns2 = append(fns2, hello)
	}
	err := Execute(fns, 10)
	if err != nil {
		fmt.Println(err)
	}

	err2 := Execute(fns2, 10)
	if err2 != nil {
		fmt.Println(err)
	}

}
