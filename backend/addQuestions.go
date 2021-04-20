
package main

import(
	"goserver/database"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"io/ioutil"
	"fmt"
	"os"
)

var (
        db *gorm.DB
)

func main(){

	fileName:=os.Args[1]

	//OPEN DATABASE
	var err error
        db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
        if err != nil {
                panic("failed opening database")
        }
        //reading in json questions
        raw, err := ioutil.ReadFile(fileName)
        if err != nil {
                fmt.Print(err)
        }

	var questionArray []database.Question
        err = json.Unmarshal(raw, &questionArray)
        if err != nil {
                fmt.Println("error reading json: ", err)
        }
	for _,question:=range questionArray{
		database.InsertQuestion(question,db)
	}
}
