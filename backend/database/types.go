package database

import (
	"fmt"
)
var(
)

type RequestData struct {
        Op string `json:"op"`
	Topic string `json:"topic"`
        No uint  `json:"no"`
}

type TopicData struct{
	ID uint `json:"id"`
	Questions []uint `json:"questions"`
}

type Topic struct {
	ID uint
	Title	string
	Questions []uint
}

func (self Topic) String () string{
	ret:= "Topic no."+fmt.Sprint(self.ID)+": "+self.Title
	return ret
}

type Question struct {
	Title string `json:"title"`
	Solution string `json:"solution"`
	TopicTitle string `"json:"topic"`
}

type QuestionData struct {
        ID uint `json:"id"`
        Title    string `json:"title"`
        Solution string `json:"solution"`
	TopicID uint `json:"topicID"`
	difficulty uint `json:"difficulty"`
	used uint `json:"used"`
}

func (self QuestionData) String ()string {
	return "Question no."+fmt.Sprint(self.ID)+" Topic no."+fmt.Sprint(self.TopicID)+": "+self.Title+" -> "+self.Solution
}

