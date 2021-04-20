package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"errors"
	"gorm.io/gorm"
	"encoding/json"
	"io/ioutil"
)

var (
	db *gorm.DB
	indexer map[string]*TopicData
)

func Start() {
	var err error
	//reading in json file
	indexerFile,err:=ioutil.ReadFile("topics.json")
	if err!=nil{
		fmt.Println("error reading topics json")
	}
	err=json.Unmarshal(indexerFile,&indexer)
	if err!=nil{
		fmt.Println("error unmarshaling json topics")
	}




	fmt.Println("start database")
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed opening database")
	}
	db.AutoMigrate(&QuestionData{})
	var result QuestionData
	db.First(&result, 1)
	Test()
}

func GetQuestion(topic string, no uint)QuestionData{

	id:= indexer[topic].Questions[no]
	var data QuestionData
	db.First(&data,id)
	return data
}


func InsertQuestion(question Question, remotedb *gorm.DB,topics map[string]*TopicData)uint {
	if remotedb==nil{
		remotedb=db
	}
	if topics==nil{
		topics=indexer
	}
	topicTitle:=question.TopicTitle
	title:=question.Title
	fmt.Println(topicTitle)
	topic:=topics[topicTitle]
	topicID:=topic.ID

	var oldQuestion QuestionData
	result:=remotedb.Where("title=?",title).Limit(1).Find(oldQuestion)
	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		println("Topic "+title+" allready exists in database")
		return 0
	}else{
		newQuestion:=&QuestionData{Title:title,Solution:question.Solution,TopicID:topicID}
		remotedb.Create(newQuestion)

		return newQuestion.ID
	}
}

func Test() {
	fmt.Println("testing database")
	fmt.Println("found Topics:")
	for topic, _:=range indexer{
		fmt.Println(topic)
		for i:=uint(0);i<3;i++{

			question:=GetQuestion(topic,i)
			fmt.Println(question.Title)
		}
	}
}
