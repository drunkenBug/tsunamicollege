package main

import (
	"net/http"

)


func main(){
	println("frontend started")

	srv:=http.FileServer(http.Dir("./dist"))

	err := http.ListenAndServe(":8080",srv)
	if err != nil {
		panic(err)
	}
}
