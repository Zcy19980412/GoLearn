package main

import (
	"fmt"
	"time"
)

func main() {
	//testChannel()
	//testTimer()
	//testTicker()
	testClose()
}

func testClose() {
	c := make(chan int)
	go func() {
		for {
			select {
			case <-c:
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("c" + time.Now().Format("2006-01-02 15:04:05"))
			}
		}
	}()

	time.Sleep(10 * time.Second)
	//when channel close, all <-channel receive {}
	close(c)
	time.Sleep(3 * time.Second)
	fmt.Println("c" + time.Now().Format("2006-01-02 15:04:05"))

}

func testTimer() {
	fmt.Println(time.Now())
	timer := time.NewTimer(1 * time.Second)
	a := <-timer.C
	fmt.Println(a)
	fmt.Println(time.Now())

}

// 间隔任务
func testTicker() {
	//tick per 5s
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	for {
	}

}

func testChannel() {
	var priority = make(chan int)
	//go leaderOneWay(priority)
	//go workerOneWay(priority)

	go leaderBothWay(priority)
	go workerBothWay(priority)

	for {
	}

}

func printer(s string) {
	for _, value := range s {
		fmt.Printf("%c", value)
		time.Sleep(500 * time.Millisecond)
	}
}

// bothWay channel
func leaderBothWay(priority chan int) {
	printer("leader")
	priority <- 1
	<-priority
	fmt.Println("leader close")
	close(priority)
}

func workerBothWay(priority chan int) {
	<-priority
	printer("worker")
	fmt.Println("\nworker done")
	priority <- 1
}

// one-way channel
func leaderOneWay(priority chan<- int) {
	printer("leader")
	priority <- 1
	close(priority)
}

func workerOneWay(priority <-chan int) {
	<-priority
	printer("worker")
}
