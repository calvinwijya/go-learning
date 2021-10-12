package main
import "fmt"

func main(){
	i ,j :=1,2
	fmt.Println(i,j)

	p:=&i
	fmt.Println(*p)

	*p =21
	fmt.Println(*p)
	
}