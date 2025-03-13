package main

import (
	"fmt"
	"time"
)

func main() {
	testChannel()
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
