package main

import "fmt"

func main(){
	var name1= "calvin"
	var name2= "calvin"

	var result bool = name1 == name2
	fmt.Println(result)

	var number1 = 100
	var number2 = 200

	fmt.Println(number1 < number2)
	fmt.Println(number1 == number2)
	fmt.Println(number1 != number2)
	fmt.Println(number1 > number2)
}