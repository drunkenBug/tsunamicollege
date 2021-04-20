package main

import(
	"os"
	"goserver/database"
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
        "gorm.io/gorm"
        "io/ioutil"
)

var (
        db *gorm.DB
)

func main(){


	//OPEN DATABASE
	var err error
        db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
        if err != nil {
                panic("failed opening database")
        }
	//delete all
	db.Where("1=1").Unscoped().Delete(&database.QuestionData{})

	//read topics 
	topicsRaw, err := ioutil.ReadFile("topics.json")
	fmt.Println(string(topicsRaw))
	if err != nil{

		fmt.Println("error reading topics")
	}
	var topics map[string]*database.TopicData
	json.Unmarshal(topicsRaw,&topics)
	fmt.Println(topics)

        //reading in json questions
	questionData, err := ioutil.ReadFile("questions.json")
        if err != nil {
                fmt.Print(err)
        }

	var questions []database.Question
        err = json.Unmarshal(questionData,&questions)
        if err != nil {
                fmt.Println("error reading json: ", err)
        }
	db.AutoMigrate(&database.QuestionData{})

	for _,question :=range questions{
		id:=database.InsertQuestion(question,db,topics)
		if id>0{
			topics[question.TopicTitle].Questions=append(topics[question.TopicTitle].Questions,id)
		}
	}
	os.Remove("topics.json")

	raw,_:=json.Marshal(topics)
	file,_:=os.Create("topics.json")
	fmt.Fprint(file,string(raw))
}
