package main

import "fmt"

var input = []int{
	5,1,2,3,4,
}

func main(){

	for i:=4;i>=0;i--{		
		for j:=4;j>=0;j--{
			fmt.Println(input[i],input[j])
			if input[i] < input[j]{
				fmt.Println("swap")
				temp := input[i]
				input[i]=input[j]
				input[j]=temp
				fmt.Println(input)

				
			}
			
		}		
	}

	//var angka = [4]int{
	//	1,
	//	2,1
	//	3,
	//	4,
	//}
	//ganti := 0
	//fmt.Println(angka)
	//fmt.Println(len(angka))
	//fmt.Println(angka[2])
	//fmt.Println(angka[1]=2)
	
}