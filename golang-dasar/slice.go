package main
import "fmt"

func main(){
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 =months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[0:5]
	fmt.Println(slice2)

	var slice3 = append(slice2,"diganti")
	fmt.Println(slice3)
	fmt.Println(months)
	fmt.Println(slice2)
	slice2[1]="ini juga"
	fmt.Println(slice3)
}