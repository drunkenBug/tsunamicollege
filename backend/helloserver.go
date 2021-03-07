package main

import (
	"net/http"
)

func process(w http.ResponseWriter, request *http.Request){
	println("hadlning new request")
	w.Write([]byte("hello server"))
}

var (
	port=":3080"
)
func main() {
	println("serving on port",port)
	http.HandleFunc("/",process)
	http.ListenAndServe(port, nil)
}
