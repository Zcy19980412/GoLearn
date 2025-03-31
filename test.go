package main

import (
	"context"
	"fmt"
	"sync"
)

type Config struct {
	Name string
	Port int
}

var waitgroup = sync.WaitGroup{}

func main() {
	background := context.Background()
	waitgroup.Add(1)

	value := context.WithValue(background, "1", Config{
		Name: "test",
		Port: 1234,
	})

	go printValue(value)

	waitgroup.Wait()

}

func printValue(v context.Context) {
	fmt.Println(v.Value("1"))
	waitgroup.Done()
}
