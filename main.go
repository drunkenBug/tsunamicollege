package main

import (
	"fmt"
	"net/http"

)


func main() {
	srv:=http.FileServer(http.Dir("./dist"))

	port:="8080"
	err := http.ListenAndServe(fmt.Sprintf(":%s", port),srv)
	if err != nil {
		panic(err)
	}
}
