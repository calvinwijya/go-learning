package main
import "fmt"

func hey (value int) interface{}{
	if value == 1{
		return 1
	}else if value == 2{
		return true
	}else{
		return "salah"
	}
}

func main(){
	var nilai interface{} = hey(1)
	fmt.Println(nilai)
}