package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("GET /goodbye", handleGoodbye)
	mux.HandleFunc("GET /goodbye/{name}", handleGoodbye)

	if err := http.ListenAndServe(":8084", mux); err!=nil {
		panic(err)
	}
}

func helloHandler(res http.ResponseWriter, req *http.Request) {

	query:=req.URL.Query()
	name:= query.Get("name")
	if name == "" {
		name = "Indigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func handleGoodbye(res http.ResponseWriter, req *http.Request) {
	name:=req.PathValue("name")
	if name == "" {
		name = "Indigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
