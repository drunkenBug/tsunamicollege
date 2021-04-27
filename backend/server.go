package main

import (
	"encoding/json"
	"fmt"
	"goserver/database"
	"net/http"
)

func getNew(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var data database.RequestNew
	err:=decoder.Decode(&data)
	if err!=nil{
		fmt.Println("json parsing error ",err)
	}

	fmt.Println("receved history:",data.History)

	var response database.QuestionData=database.GetNewQuestion(data.Topic, data.Rating, data.History)
	fmt.Println("sending id ",response.ID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}


func main() {
	database.Start()
	http.HandleFunc("/api/new", getNew)
	http.ListenAndServe(":3080", nil)
}
