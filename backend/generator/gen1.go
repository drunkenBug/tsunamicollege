package main

import (
	"fmt"
	"strings"
)

func main(){
	grand:="part part"
	splits:=strings.Split(grand," ")
	fmt.Println(splits[0])
}

