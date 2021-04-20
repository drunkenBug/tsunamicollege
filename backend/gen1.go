package main

import (
	"fmt"
	"encoding/json"
	db "goserver/database"
)

type operator struct {
	symbol string
	title string
	apply func(int,int) int
}

func generateQuestions(X []int, Y []int, op operator)[]db.Question{
	arr:=[]db.Question{}
	for _,x:=range X{
		for _,y:=range Y{
			title:="was ist "+fmt.Sprint(x)+op.symbol+fmt.Sprint(y)
			solution:=fmt.Sprint(op.apply(x,y))
			next:=db.Question{Title:title,Solution:solution,TopicTitle:op.title}
			arr=append(arr,next)
		}
	}
	return arr
}

func generateDivisions(Y []int, Z []int, op operator)[]db.Question{
	arr:=[]db.Question{}
	for _,y:=range Y{
		for _,z:=range Z{
			//idea: x:y=z
			x:=y*z
			title:="was ist " + fmt.Sprint(x)+op.symbol+fmt.Sprint(y)
			solution:=fmt.Sprint(z)
			next:=db.Question{Title:title,Solution:solution,TopicTitle:op.title}
			arr=append(arr,next)
		}
	}
	return arr
}

func main(){
	plus:=operator{title:"Plus",symbol:"+",apply:func(a int, b int)int{
		return a+b
	}}
	minus:=operator{title:"Minus",symbol:"-",apply:func(a int, b int)int{
		return a-b
	}}
	mal := operator {title:"Ein mal eins",symbol:"·",apply:func(a int, b int)int {
		return a*b
	}}
	geteilt := operator {title:"Geteilt",symbol:":",apply:func(a int, b int)int{
		return a/b
	}}

	var oneten=[]int {1,2,3,4,5,6,7,8,9,10}
	var twoten=[]int {2,3,4,5,6,7,8,9,10}

	var questions []db.Question

	questions=append(questions,generateQuestions(oneten,oneten,plus)...)
	questions=append(questions,generateQuestions(oneten,oneten,minus)...)
	questions=append(questions,generateQuestions(oneten,oneten,mal)...)
	questions=append(questions,generateDivisions(twoten,twoten,geteilt)...)


	raw,_:=json.Marshal(questions)
	fmt.Println(string(raw))

}

