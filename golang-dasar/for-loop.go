package main

import "fmt"

func main(){
	counter := 1

	for counter <= 10{
		fmt.Println("perulangan ke ", counter)
		counter++
	}
	
	
	for counter:= 1;counter <= 10;counter++{
		fmt.Println("perulangan ke", counter)
	}


	slice := []string{"calvin","wijaya","acin","vio","abem"}

	for i := 0 ; i < len(slice) ; i++{
		fmt.Println(slice[i])
	}

	for _, value := range slice{
		fmt.Println(value)
	}

	

	
}