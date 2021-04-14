package main

import(
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
	db.Where("1=1").Unscoped().Delete(&database.Question{})
	db.Where("1=1").Unscoped().Delete(&database.Topic{})

        //reading in json questions
        questionData, err := ioutil.ReadFile("questions.json")
        if err != nil {
                fmt.Print(err)
        }

	var questionMap map[string][]database.Question
        err = json.Unmarshal(questionData, &questionMap)
        if err != nil {
                fmt.Println("error reading json: ", err)
        }
	db.AutoMigrate(&database.Topic{})
	db.AutoMigrate(&database.Question{})
	
	for topicTitle,questions :=range questionMap{
		topic:=database.Topic{Title:topicTitle,Questions:questions}
		db.Create(&topic)
	}
	fmt.Println("testing...")
	database.DBprint(db)
	topic:="Plus"
	no:=1
	fmt.Println("find where topic="+topic+" id = "+fmt.Sprint(no))
	var t1 database.Topic
	db.Where("title = ?",topic).First(&t1)
	//var q1 database.Question
	var qs []database.Question
	db.Model(&t1).Association("Questions").Find(&qs)

	fmt.Println(qs[no])

	fmt.Println("\ncreated Database printing results:")
}
