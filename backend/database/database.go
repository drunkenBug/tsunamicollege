package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
)

var (
	db      *gorm.DB
	indexer map[string]*TopicData
)

func Start() {
	var err error
	//reading in json file
	indexerFile, err := ioutil.ReadFile("topics.json")
	if err != nil {
		fmt.Println("error reading topics json")
	}
	err = json.Unmarshal(indexerFile, &indexer)
	if err != nil {
		fmt.Println("error unmarshaling json topics")
		panic(err)
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

func InsertQuestion(question Question, remotedb *gorm.DB, topics map[string]*TopicData) int {
	if remotedb == nil {
		remotedb = db
	}
	if topics == nil {
		topics = indexer
	}
	topicTitle := question.TopicTitle
	title := question.Title
	fmt.Println(topicTitle)
	topic := topics[topicTitle]
	topicID := topic.ID

	var oldQuestion QuestionData
	result := remotedb.Where("title=?", title).Limit(1).Find(oldQuestion)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		println("Topic " + title + " allready exists in database")
		return 0
	} else {
		newQuestion := &QuestionData{Title: title, Solution: question.Solution, Rating: question.Rating, TopicID: topicID}
		remotedb.Create(newQuestion)

		newEntry :=QuestionNote{ID: newQuestion.ID, Rating: newQuestion.Rating}
		topic.Questions = append(topic.Questions, QuestionNote{ID: newQuestion.ID, Rating: newQuestion.Rating})
		index := len(topic.Questions) - 2
		for index >= 0 && topic.Questions[index].Rating > newEntry.Rating {
			topic.Questions[index+1] = topic.Questions[index]
			index -= 1
		}
		topic.Questions[index+1] = newEntry
		//fmt.Println("insert question ",newQuestion.ID,"at rating ",newQuestion.Rating)
		return newQuestion.ID
	}
}

//finds the index of a question with given rating in TopicData
func getNumberByRating(topicData TopicData, rating int) int {
	var questions []QuestionNote
	questions = topicData.Questions
	var min, max int

	min = 0
	max = len(questions) - 1
	if questions[min].Rating >= rating {
		return min
	}
	if questions[max].Rating <= rating {
		return max
	}
	var next int
	var t float32
	for {
		t = float32(rating - questions[min].Rating) /float32(questions[max].Rating - questions[min].Rating)
		next = int(float32(min)*(1-t) + float32(max)*t)
		if questions[next].Rating > rating {
			max = next
		} else if questions[next].Rating < rating {
			min = next
		}
		if next == max || next == min || questions[next].Rating == rating {
			break
		}
	}
	return next
}

func find(ls []int, el int)bool{
	for _,found := range ls{
		if found == el{
			return true
		}
	}
	fmt.Println("not found",el,ls)
	return false
}

func GetNewQuestion(topic string, rating int, history []int)QuestionData{

	var no int
	topicData:=*(indexer[topic])
	var questions []QuestionNote
	questions=topicData.Questions

	no=getNumberByRating(topicData,rating)
	var id int = questions[no].ID
	var last int = len(questions)-1
	for {
		if no ==last {
			break
		}
		if !find(history, id){
			break
		}
		no++
		id = questions[no].ID
	}
	var questionData QuestionData
	db.Find(&questionData, id)
	return questionData
}

func Test() {
	fmt.Println("testing database")
	fmt.Println("indexer 'Plus':")
	fmt.Println(*indexer["Plus"])

}
