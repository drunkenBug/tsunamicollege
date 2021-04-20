package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	"goserver/database"
)

func process(w http.ResponseWriter, request *http.Request){
	//fmt.Println("processing request")
	decoder:=json.NewDecoder(request.Body)

	var init database.RequestData
	decoder.Decode(&init)
	fmt.Println(init)

	var response database.QuestionData
	response=database.GetQuestion(init.Topic,init.No)

	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	if err:= json.NewEncoder(w).Encode(response);err!=nil{
		panic(err)
	}
}


func processSet(w http.ResponseWriter, request *http.Request){
	decoder:=json.NewDecoder(request.Body)
	var inserter database.Question
	decoder.Decode(&inserter)

	database.InsertQuestion(inserter,nil,nil)

}

func main() {
	database.Start()
	http.HandleFunc("/api/getQuestion",process)
	http.HandleFunc("/api/setQuestion",processSet)
	http.ListenAndServe(":3080", nil)
}
