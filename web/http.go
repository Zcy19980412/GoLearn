package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"net/http"
	"os"
)

func main() {

	//testNet()
	//testHttp()
	testMysql()

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

func testMysql() {
	open, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/todolist")
	errFunc(err, "open")
	defer open.Close()

	query, err := open.Query("select * from habit limit 10")
	errFunc(err, "query")
	defer query.Close()
	for query.Next() {
		columns, _ := query.Columns()
		var id, time, time1, userid, gapdays, name, desc, rate, imrate string
		fmt.Println(columns)
		err := query.Scan(&id, &time, &time, &userid, &gapdays, &name, &desc, &rate, &imrate)
		errFunc(err, "scan")
		fmt.Println(id, time, time1, userid, gapdays, name, desc, rate, imrate)
	}
}
