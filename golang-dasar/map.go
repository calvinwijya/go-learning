package main

import "fmt"

func main(){
	person := map[string]string{
		"nama":"calvin",
		"hobby":"berenang",
	}
	fmt.Println(person)
	fmt.Println("nama")
	fmt.Println("hobby")

	person["title"]="programmer"

	fmt.Println(person)

	pacar := make(map[string]string)
		pacar["pertama"]="christine"
		pacar["kedua"]="ga ada"
		pacar["ketiga"]="jombs"

		fmt.Println(pacar)

		delete(pacar,"pertama")

		fmt.Println(pacar,)
	

}
