package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
)

func init() {
	log.SetOutput(os.Stdout)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func logger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		log.Printf("Handle function - %s", name)
		f(w, r)
	}
}

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}

	http.HandleFunc("/hello", logger(handler))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "index")
	})

	server.ListenAndServe()
}
