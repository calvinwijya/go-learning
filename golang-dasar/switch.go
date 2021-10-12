package main

import "fmt"

func main(){
	name := "calvin wijaya"

	switch name{
		case "cavin wijaya":
			fmt.Println("botak")
	
		case "fernando":
			fmt.Println("acin")
	
		default:
			fmt.Println("nda kenal")
}


		switch length := len(name) ; length > 8{

		case true :
			fmt.Println("panjang bet")
		case false:
			fmt.Println("oke pas")	
}

		length := len(name)
		switch{
		case length > 10:
			fmt.Println("pnjng")
		case length > 5:
			fmt.Println("okela")
		default:
			fmt.Println("ws")
		}


}