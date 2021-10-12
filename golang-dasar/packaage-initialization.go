package main

import (
	"golang-dasar/database"
	"fmt"
)
	
func main(){
	result := database.GetDataBase()
	fmt.Println(result)
}