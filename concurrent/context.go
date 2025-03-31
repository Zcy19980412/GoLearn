package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	testContextWithValue()

	//context作为一个信号接收器，接收（select）其他go程的done信号
	//testContextWithCancel()

	//和Deadline一样，超过deadline直接close Context.Done()
	//testContextTimeOut()
}
func testContextWithValue() {
	background := context.Background()
	withValue := context.WithValue(background, "key", "value")
	go func() {
		fmt.Println(withValue)
	}()
	time.Sleep(1 * time.Second)
}
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
func testContextTimeOut() {
	backGround := context.Background()
	timeout, cancel := context.WithTimeout(backGround, 10*time.Second)

	go func(timeout context.Context) {
		for {
			select {
			case c := <-timeout.Done():
				fmt.Println(c)
				return
			default:
				fmt.Println("hello world")
				time.Sleep(1 * time.Second)
			}
		}
	}(timeout)
	time.Sleep(13 * time.Second)
	cancel()
	for {

	}

}
