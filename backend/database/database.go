package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Start() {
	var err error
	//reading in json file

	fmt.Println("start database")
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed opening database")
	}
	db.AutoMigrate(&Question{})
	var result Question
	db.First(&result, 1)
	//Test()
}

func GetById(id uint)Question{
	var question Question
	db.First(&question,id)
	return question
}

func InsertQuestion(data InsertData) Question{
	fmt.Println("insert question")
	question:=Question{Title: data.Title, Solution:data.Solution}
	db.Create(&question)
	fmt.Println("inser ID: ",question.ID)
	return question
}

func Test() {
	fmt.Println("test")
	var questions []Question
	result:=db.Find(&questions)
	fmt.Println("result: ",result.RowsAffected," found")
	for _,q :=range questions{
		fmt.Println("question: ",q)
	}
	for i:=uint(1);i<4;i+=1{
		question:=GetById(i)
		fmt.Println(question.Title)
	}
}
