package main

import "fmt"

func main (){
	var nilaiAkhir = 90
	var absensi = 81

	var lulusnilaiAkhir = nilaiAkhir > 90
	var lulusAbsensi = absensi > 80

	var lulus bool = lulusnilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
func calculateRectangleArea(length float64, width float64) float64 {
    return length * width