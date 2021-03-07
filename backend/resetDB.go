package main

import(
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
        "gorm.io/gorm"
        "io/ioutil"
)
type Question struct {
        Id uint `json:"id"`
        Title    string `json:"title"`
        Solution string `json:"solution"`
}

var (
        db *gorm.DB
)

func main(){
	var err error
        //reading in json file
        fileData, err := ioutil.ReadFile("questions.json")
        if err != nil {
                fmt.Print(err)
        }

        var questions []Question
        err = json.Unmarshal(fileData, &questions)
        if err != nil {
                fmt.Println("error reading json: ", err)
        }


        fmt.Println("start database")

	//OPEN DATABASE
        db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//delete all
	db.Where("1=1").Unscoped().Delete(&Question{})

        if err != nil {
                panic("failed opening database")
        }
        fmt.Println("opened database")
        db.AutoMigrate(&Question{})

        for _,questionData :=range questions{
                fmt.Println("adding question: ",questionData.Title)
                db.Create(&Question{Title: questionData.Title,Solution: questionData.Solution})
        }
}
