package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {

	//testNet()
	testHttp()

}

func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}

func testNet() {
	listen, err := net.Listen("tcp", ":8080")
	errFunc(err, "listen")
	defer listen.Close()

	accept, err := listen.Accept()
	errFunc(err, "accept")
	defer accept.Close()

	bytes := make([]byte, 1024)
	_, err = accept.Read(bytes)
	errFunc(err, "read")
	fmt.Println(string(bytes))
	_, err = accept.Write([]byte("hello world"))
	errFunc(err, "write")
}

func testHttp() {
	//http:  HandleFunc  ListenAndServe
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080", nil)
}
