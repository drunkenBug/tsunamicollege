package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"goserver/database"
	"io/ioutil"
	"os"
)

var (
	db *gorm.DB
)

func main() {

	//OPEN DATABASE
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed opening database")
	}
	//delete all
	db.Where("1=1").Unscoped().Delete(&database.QuestionData{})

	//read topics
	topicsRaw, err := ioutil.ReadFile("emptyTopics.json")
	fmt.Println(string(topicsRaw))
	if err != nil {

		fmt.Println("error reading topics")
	}
	var topics map[string]*database.TopicData
	err=json.Unmarshal(topicsRaw, &topics)
	if err!=nil{
		fmt.Println("error parsing json topics")
		panic (err)
	}
	//fmt.Println(topics)

	//reading in json questions
	questionData, err := ioutil.ReadFile("questions.json")
	if err != nil {
		fmt.Print(err)
	}

	var questions []database.Question
	err = json.Unmarshal(questionData, &questions)
	if err != nil {
		fmt.Println("error reading json: ", err)
	}
	db.AutoMigrate(&database.QuestionData{})

	for _, question := range questions {
		//fmt.Println("q: ",question.Title,"r: ",question.Rating)
		database.InsertQuestion(question, db, topics)

	}
	os.Remove("topics.json")

	raw, _ := json.Marshal(topics)
	file, _ := os.Create("topics.json")
	fmt.Fprint(file, string(raw))
}
