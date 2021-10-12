package main

import "fmt"

type HasName interface{
	getName() string
}
func sayHello(hasname HasName){
	fmt.Println("hello",hasname.getName())
}


type Calvin struct{
	Name string
}
func (calvin Calvin) getName() string{
	return calvin.Name
}

func main(){
	var person Calvin
	person.Name="calvin"
	sayHello(person)
}