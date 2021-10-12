package main


import "fmt"

type Address struct{
	City,Province,Country string
}

func ChangeCountryTo(address *Address){
	Address.Country="indonesias"
}

func main (){
	address1:= Address{"batam","kepri","indonesia"}
	address2:= &address1
	address3:= &address2

	address2.City="tanjung pinang"

	

	*address2 = Address{"jakarta","jawa","indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(*address3)
	
	alamat := Address{"aceh","sumut",""}

	ChangeCountryTo(&alamat)

}