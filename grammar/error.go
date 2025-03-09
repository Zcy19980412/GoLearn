package main

import (
	"errors"
	"fmt"
)

func main() {
	//自己发现的错误用error来处理，否则使用defer+recover 处理panic
	//testError
	//testError()

	//testPanic
	testPanic()
}

func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Divide(10, 0)
	fmt.Println("end")
}

func testError() {
	//i := Divide(10, 0)
	i, err := DivideErrorHandled(10, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	//panic("123")
	fmt.Println("end")
}

func Divide(num1 int, num2 int) int {
	return num1 / num2
}

func DivideErrorHandled(num1 int, num2 int) (result int, err error) {
	if num2 == 0 {
		err = errors.New("divide by zero")
		return
	}
	result = num1 / num2
	return
}
