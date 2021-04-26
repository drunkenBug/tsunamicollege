package main

import (
	"encoding/json"
	"fmt"
	"goserver/database"
	"net/http"
)

func process(w http.ResponseWriter, request *http.Request) {
	//fmt.Println("processing request")
	decoder := json.NewDecoder(request.Body)

	var init database.RequestData
	decoder.Decode(&init)
	fmt.Println(init)

	var response interface{}
	if init.Op == "GET" {
		response = database.GetQuestion(init.Topic, init.No)
	} else if init.Op == "GETINFO" {
		response = database.GetTopicInformation(init.Topic)
	} else if init.Op == "GETBYRATING" {
		response = database.GetQuestionByRating(init.Topic, init.No)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func main() {
	database.Start()
	http.HandleFunc("/api/", process)
	http.ListenAndServe(":3080", nil)
}
