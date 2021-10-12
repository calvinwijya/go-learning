package main

import "fmt"

func factorialLoop (value int)int{
	result := 1
	for i := value ; i > 1 ; i--{
		result *= i
	}
	return result
}

func main(){

	loop:= factorialLoop(5)
	fmt.Println(loop)

	looping:= factorialRecursive(5)
	fmt.Println(looping)
}


func factorialRecursive(value int)int{
	if value == 1{
		return 1
	}else{
		return value * factorialRecursive(value - 1)
	}
}