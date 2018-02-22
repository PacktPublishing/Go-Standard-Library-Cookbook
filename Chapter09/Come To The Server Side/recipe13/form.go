package main

import (
	"fmt"
	"net/http"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Prior ParseForm: %v\n", req.Form)
	req.ParseForm()
	fmt.Printf("Post ParseForm: %v\n", req.Form)
	fmt.Println("Param1 is : " + req.Form.Get("param1"))
	fmt.Printf("PostForm : %v\n", req.PostForm)
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

func main() {
	s := createServer(":8080")
	fmt.Println("Server is starting...")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
