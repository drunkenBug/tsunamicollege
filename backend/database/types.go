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

type InsertData struct {
        Op string `json:"op"`
        Title string `json:"title"`
        Solution string `json:"solution"`
}
type Topic struct {
	ID uint `json:"id"`
	Title	string `json:"title"`
	Questions []Question
}

func (self Topic) String () string{
	ret:= "Topic no."+fmt.Sprint(self.ID)+": "+self.Title
	for i,_ :=range self.Questions{
		ret+=fmt.Sprint(i)
	}
	return ret

}

type Question struct {
        ID uint `json:"id"`
        Title    string `json:"title"`
        Solution string `json:"solution"`
	TopicID uint `json:"toicId"`
}

func (self Question) String ()string {
	return "Question no."+fmt.Sprint(self.ID)+" Topic no."+fmt.Sprint(self.TopicID)+": "+self.Title+" -> "+self.Solution
}
