package main

import (
	"fmt"
	"net/http"
)

type handle struct {}

func (h handle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func body(w http.ResponseWriter, r *http.Request)  {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintf(w, string(body)+"\n\n")

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
