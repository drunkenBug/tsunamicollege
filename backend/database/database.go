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
	Test()
}

func GetQuestion(topic string, no uint)Question{
	//topic:="plus"
        //no:=1
        fmt.Println("find where topic="+topic+" no = "+fmt.Sprint(no))
        var t1 Topic
        db.Where("title = ?",topic).First(&t1)
        //var q1 database.Question
        var qs []Question
        db.Model(&t1).Association("Questions").Find(&qs)

        return (qs[no-1])
}

func InsertQuestion(data InsertData) Question{
	fmt.Println("insert question")
	question:=Question{Title: data.Title, Solution:data.Solution}
	db.Create(&question)
	fmt.Println("inser ID: ",question.ID)
	return question
}

func Test() {

	fmt.Println("test database: ")
	DBprint(db)
	question:=GetQuestion("Plus",1)
	fmt.Println("retrieved smaple questin: ",question)
}
