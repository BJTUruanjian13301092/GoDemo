package main

import (
	"io"
	"net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "hello Yitong")
	io.WriteString(rw, "\n")
	io.WriteString(rw, req.URL.String())
}

func main() {
	http.HandleFunc("/", hello)  //设定访问的路径
	http.ListenAndServe(":8080", nil) //设定端口和handler
}
