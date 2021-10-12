package main



import "fmt"

type filter func(string)string


func sayHelloWithFilter(name string,filter filter){
	filteredName := filter(name)
	fmt.Println("hello",filteredName)
}

func spamFilter(name string)string{
	if name == "anjing"{
		return "..."
	}else{
		return name
	}
}

func main(){
	sayHelloWithFilter("calvin",spamFilter)
	sayHelloWithFilter("anjing",spamFilter)

	filter:= spamFilter
	sayHelloWithFilter("anjing",filter)

}