package main

import "fmt"

func hitungBiaya(jenis string, masuk, keluar int) int {
	var tarif int
	var total int
	var jamParkir int

	jamParkir = keluar - masuk

	for i := 0; i < jamParkir; i++ {

		if jenis == "motor" {
			if masuk+i < 17 {
				tarif = 4000
			} else {
				tarif = 5000
			}
		} else if jenis == "mobil" {
			if masuk+i < 17 {
				tarif = 6000
			} else {
				tarif = 7000
			}
		}

		total = total + tarif
	}

	return total
}

func main() {

	var jenis string
	var masuk, keluar int
	var totalPendapatan int
	var biaya int
	var nomor int = 1

	fmt.Println("=== Sistem Parkir Perpustakaan Cafe ===")

	for {

		fmt.Println("\nData Kendaraan ke-", nomor)

		fmt.Print("Jenis kendaraan (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Jam masuk (0-24): ")
		fmt.Scan(&masuk)

		fmt.Print("Jam keluar (0-24): ")
		fmt.Scan(&keluar)

		biaya = hitungBiaya(jenis, masuk, keluar)

		fmt.Println("Total biaya parkir:", biaya)
		fmt.Println("-----------------------------")

		totalPendapatan = totalPendapatan + biaya
		nomor++
	}

	fmt.Println("\nTotal pendapatan parkir hari ini:", totalPendapatan)
}
