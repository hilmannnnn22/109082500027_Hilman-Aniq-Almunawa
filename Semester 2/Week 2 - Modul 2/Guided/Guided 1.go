package main

import "fmt"

func main() {
	var satu string
	var dua string
	var tiga string
	var temp string

	fmt.Print("Masukkan input: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output Awal: " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output Akhir" + satu + " " + dua + " " + tiga)
}
