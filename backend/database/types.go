package database

import (
	"fmt"
)
var(
)

type RequestData struct {
        Op string `json:"op"`
	Topic string `json:"topic"`
        No int  `json:"no"`
}

type QuestionNote struct{
	ID int `json:"id"`
	Rating int `json"rating"`
}

type TopicData struct{
	ID int `json:"id"`
	Questions []QuestionNote `json:"questions"`
}

type TopicInformation struct{
	Size int `json:"size"` //how many questions are in this topic?
	Bottom int `json:"bottom"` //whats the difficulty of the easiest question
	Top int `json:"top"`//whats the difficulty of the hardest question 
}

type Topic struct {
	ID int
	Title	string
	Questions []int
}

func (self Topic) String () string{
	ret:= "Topic no."+fmt.Sprint(self.ID)+": "+self.Title
	return ret
}

type Question struct {
	Title string `json:"title"`
	Solution string `json:"solution"`
	TopicTitle string `"json:"topic"`
	Rating int `json:"rating"`
}

type QuestionData struct {
        ID int `json:"id"`
        Title    string `json:"title"`
        Solution string `json:"solution"`
	TopicID int `json:"topicID"`
	Rating int `json:"rating"`
	used int `json:"used"`
}

func (self QuestionData) String ()string {
	return "Question no."+fmt.Sprint(self.ID)+" Topic no."+fmt.Sprint(self.TopicID)+": "+self.Title+" -> "+self.Solution
}

