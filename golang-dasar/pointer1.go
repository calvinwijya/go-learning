package main


import "fmt"

func main(){
	a:= 1
	fmt.Println(a)
	fmt.Println(&a)

	var b *int 
	b=&a
	fmt.Println(b)
	fmt.Println(*b)


	aa:=5
	bb:=5
	Calc(&aa,&bb)

	c:=&aa
	d:=&c
	fmt.Println(**d)
	fmt.Println(c)
}

func Calc(a,b *int){
	fmt.Println(*a + *b)
}