package main

import (
	"context"
	"fmt"
)

func main() {
	background := context.Background()
	background2 := context.Background()

	fmt.Printf("background: %p\n", background)
	fmt.Printf("background2: %p\n", background2)

	if background == background2 {
		fmt.Println("background 和 background2 是同一个实例")
	} else {
		fmt.Println("background 和 background2 不是同一个实例")
	}
}
