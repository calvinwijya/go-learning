package main
 
import "fmt"

func main(){
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAKhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var kelulusan bool = lulusNilaiAKhir && lulusAbsensi
	var lulus bool = lulusNilaiAKhir || lulusAbsensi
	fmt.Println(kelulusan)
	fmt.Println(lulus)
}