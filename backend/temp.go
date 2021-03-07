package main

type data struct{
	ID uint
}

type question struct{
	ID uint
	title string
}

type theme struct{
	ID uint
	name string
}

func (member data) getID()uint{
	return member.ID
}

func getID(d data)uint{
	return d.ID
}

func main(){
	q:=question{}
	q.ID=12
	q.title="whats 4+4"
	id:=getID(q)
	if id==12{
		println("test success")
	}else{
		println("test faied")
	}
}
