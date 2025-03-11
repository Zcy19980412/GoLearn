package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//testCreate()
	//testWrite()
	testRead()
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

	//seek
	str = "seek"
	strBytes = []byte(str)
	fileContentEndNum, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteAt(strBytes, fileContentEndNum)
	if err != nil {
		return
	}

}

func testRead() {
	openFile, err := os.OpenFile("D://test.txt", os.O_APPEND|os.O_RDONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	defer openFile.Close()
	var contents = make([]byte, 10)
	//startNum, err := openFile.Seek(0, io.SeekStart)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//_, err = openFile.ReadAt(contents, startNum)
	//if err != nil && err != io.EOF {
	//	fmt.Println(err)
	//}
	//File的光标默认在index = 0 处
	for {
		_, err := openFile.Read(contents)
		if err == io.EOF {
			break
		}
		fmt.Println(string(contents))
	}
	fmt.Println(string(contents))

}
