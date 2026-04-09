package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
}

func biayaPerHari(jumlahMhs int) int {
	makan := 35000 * 2
	penginapan := 250000
	uangSaku := 300000
	total := makan + penginapan + uangSaku
	return total * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hari := tanggunganHari(lamaPerjalanan, tujuan)
	biayaHarian := biayaPerHari(jumlahMhs)

	if tujuan == "mancanegara" {
		biayaHarian = int(float64(biayaHarian) * 1.5)
	}

	*totalBiaya = float64(hari * biayaHarian)
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	fmt.Println()

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("Biay perjalanan yang harus dikeluarkan Tel-U: %.0f\n", biaya)
}
