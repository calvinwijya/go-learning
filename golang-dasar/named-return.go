package main

import "fmt"

func getCompleteName()(firstName,middleName string,lastName int){
	firstName="calvin"
	middleName="wijaya"
	lastName=31

	return firstName,middleName,lastName
}

func main(){
	firstName,middleName,lastName := getCompleteName()
	fmt.Println(firstName,middleName,lastName)
}