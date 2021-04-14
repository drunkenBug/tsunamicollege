package database

import (
	"fmt"
	"gorm.io/gorm"
)

func DBprint(db *gorm.DB){
	var questions []Question
	db.Find(&questions)
	println("found questions in database: ")
	for _,q:=range(questions){
		fmt.Println(q)
	}
	var topics []Topic
	db.Find(&topics)
	println("found topics in database: ")
	for _,t:=range(topics){
		fmt.Println(t)
	}
}
