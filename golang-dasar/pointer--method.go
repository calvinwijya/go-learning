package main
import "fmt"


type Me struct{
	Name string
}

func (me *Me) Status (){
	me.Name ="what's up "+ me.Name
}

func main(){
	calvin:= Me{"calvin"}

	calvin.Status()
	fmt.Println(calvin.Name)
}