package main

import "fmt"

type Customers struct{
	Name,Live string
	Age int
}

func (customers Customers) sayHello(name string){
	fmt.Println("hello",name,"my name is",customers.Name)
}

func main(){
	calvin := Customers{
		Name:"calvin",
		Live:"batam",
		Age:19,
	}

	calvin.sayHello("acin")
}