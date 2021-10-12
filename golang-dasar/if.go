package main

import"fmt"

func main(){
	name := "why"

	if name == "calvin"{
		fmt.Println("cibai")
	}else if name =="why"{
		fmt.Println("betul")
	}else{
		fmt.Println("salah")
	}


	
	if length := len(name) ; length > 5{
		fmt.Println("pas")
	}else{
		fmt.Println("gak pas")
	}
}