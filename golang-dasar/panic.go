package main

import "fmt"

func endApp(){
	message := recover()
	if message != nil{
		fmt.Println("terjadi error dengan message: ", message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("eERROR")
	}
}


func main(){
	runApp(false)
}