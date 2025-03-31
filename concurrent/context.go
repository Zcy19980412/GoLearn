package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	//testContextWithValue()
	testContextWithCancel()

}

// context作为一个信号接收器，接收（select）其他go程的done信号
func testContextWithCancel() {
	backGround := context.Background()
	withCancel, cancelFunc := context.WithCancel(backGround)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)

	go func(withCancel context.Context, waitGroup *sync.WaitGroup) {
		for {
			select {
			case c := <-withCancel.Done():
				fmt.Println(c)
				fmt.Println("1")
				waitGroup.Done()
				fmt.Println("2")
				return
			default:
				fmt.Println("hello world")
			}
		}

	}(withCancel, waitGroup)

	go func(withCancel context.Context, waitGroup *sync.WaitGroup) {
		for {
			select {
			case c := <-withCancel.Done():
				fmt.Println(c)
				fmt.Println("11")
				waitGroup.Done()
				fmt.Println("22")
				return
			default:
				fmt.Println("hello world")
			}
		}

	}(withCancel, waitGroup)

	time.Sleep(1 * time.Second)
	cancelFunc()
	time.Sleep(10 * time.Second)
	fmt.Println("3")
	waitGroup.Wait()
}
