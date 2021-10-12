package main

import "fmt"

func getGoodBye(name string)string{
	return "goodbye " + name
}

func main(){
	goodbye := getGoodBye("calvin")
	fmt.Println(goodbye)
}