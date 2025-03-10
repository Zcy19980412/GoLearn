package main

import (
	"fmt"
	"os"
)

func main() {
	//testCreate()
	testWrite()
}

func testCreate() {
	create, err := os.Create("D://test.txt")
	if err != nil {
		err.Error()
		fmt.Println(err)
	}
	defer create.Close()

}

func testWrite() {
	file, err := os.Create("D://test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
	}

	str := "Hello World"
	strBytes := []byte(str)
	_, err = file.Write(strBytes)
	if err != nil {
		fmt.Println(err)
	}

}
