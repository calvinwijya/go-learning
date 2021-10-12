package main

import "fmt"

type Customer struct {
	Name, Address string
	PhoneNumber   int
}

func (customer Customer) sayHi(name string){
	fmt.Println("hello", name,"my name is ", customer.Name)
}
func (c Customer) sayBoo(){
	fmt.Println("BOOOOOOOOO from" , c.Name)
}

func main(){
	var calvin  Customer
	calvin.Name="calvin"
	calvin.Address="taman baloi mas"
	calvin.PhoneNumber=89789786
	
	calvin.sayHi("acin")	

	calvin.sayBoo()

}