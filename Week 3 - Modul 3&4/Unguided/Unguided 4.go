package main

import "fmt"

func hitungPersegi(sisi int) {
	var luas, keliling int

	luas = sisi * sisi
	keliling = 4 * sisi

	fmt.Println("Luas persegi :", luas)
	fmt.Println("Keliling persegi :", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	var luas, keliling int

	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	fmt.Println("Luas persegi panjang :", luas)
	fmt.Println("Keliling persegi panjang :", keliling)
}

func hitungLingkaran(jarijari float64) {
	var luas, keliling float64
	var pi float64 = 3.14

	luas = pi * jarijari * jarijari
	keliling = 2 * pi * jarijari

	fmt.Println("Luas lingkaran :", luas)
	fmt.Println("Keliling lingkaran :", keliling)
}

func main() {

	var pilihan int

	fmt.Println("=== PROGRAM HITUNG BANGUN DATAR ===")
	fmt.Println("1. Persegi")
	fmt.Println("2. Persegi Panjang")
	fmt.Println("3. Lingkaran")

	fmt.Print("Pilih menu : ")
	fmt.Scan(&pilihan)

	switch pilihan {

	case 1:
		var sisi int
		fmt.Print("Masukkan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)

	case 2:
		var panjang, lebar int
		fmt.Print("Masukkan panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)

	case 3:
		var r float64
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&r)
		hitungLingkaran(r)

	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}
