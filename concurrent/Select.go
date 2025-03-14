package main

import (
	"fmt"
	"time"
)

func main() {
	//testSelect()
	testTimeout()
}

/*
select 同时监听多个channel的状态，在channel满足条件时，执行channel的读写操作。 在go程 1：n的情况下，赋予主go程并发处理go程的能力。
select 的确像 CPU 的并发调度机制。它通过等待多个 channel 的状态变化，智能地选择执行某个任务，从而优化了 Go 程序的并发执行效率。
这个类比是非常准确的，帮助理解了 select 的作用：它赋予了协程多路复用的能力，并根据多个通道的状态“调度”并发任务，最终提升了效率。
*/
func testSelect() { // 创建两个无缓冲通道
	chan1 := make(chan int)
	chan2 := make(chan int)

	// 启动 goroutine，向 chan1 发送 5 个整数，并在每次发送后等待 1 秒
	go func() {
		for i := 0; i < 5; i++ {
			chan1 <- i
			time.Sleep(time.Second)
		}
		close(chan1) // 发送完数据后关闭通道
	}()

	// 启动 goroutine，向 chan2 发送 5 个整数，并在每次发送后等待 1 秒
	go func() {
		for i := 45; i < 50; i++ {
			chan2 <- i
			time.Sleep(time.Second)
		}
		close(chan2) // 发送完数据后关闭通道
	}()

	// 传统的 `for i := range chan1` 方式是串行的：
	// - 先遍历 chan1，直到它关闭
	// - 再遍历 chan2，直到它关闭
	// - 不能交替处理两个通道的数据
	//
	// for i := range chan1 {
	//     fmt.Println(i)
	// }
	// for i := range chan2 {
	//     fmt.Println(i)
	// }

	// 使用 `select` 让主 goroutine 可以异步监听多个 `channel`
	// - `select` 让主 goroutine 同时监听 `chan1` 和 `chan2` 的数据
	// - `chan1` 和 `chan2` 的数据可以交错输出（谁先有数据就处理谁）
	// - `ok` 用于检测通道是否关闭，关闭后我们用 `chan1 = nil` 避免 select 继续监听
	for {
		select {
		case v1, ok := <-chan1:
			if !ok {
				// chan1 关闭后，将其置为 nil，避免 select 继续监听它
				chan1 = nil
			} else {
				fmt.Println("chan1 :", v1)
			}

		case v2, ok := <-chan2:
			if !ok {
				// chan2 关闭后，将其置为 nil，避免 select 继续监听它
				chan2 = nil
			} else {
				fmt.Println("chan2 :", v2)
			}
		}

		// 当两个通道都关闭时，退出循环
		if chan1 == nil && chan2 == nil {
			break
		}
	}

}

func testTimeout() {
	c := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		c <- 1
	}()

	select {
	case i := <-c:
		fmt.Println("receive:", i)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}

}
