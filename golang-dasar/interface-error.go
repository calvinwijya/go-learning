package main

import (
	"fmt"
	"errors"
)

func Pembagian (nilai int, pembagi int)(int,error){
	if pembagi == 0{
		return 0, errors.New("error karena terbagi karena nol")
	}else{
		result := nilai / pembagi
		return result, nil
	}
}

func main(){
	hasil,err:=Pembagian(10,0)
	if err == nil{
		fmt.Println("hasil",hasil)
	}else{
		fmt.Println("error",err.Error())
	}
}