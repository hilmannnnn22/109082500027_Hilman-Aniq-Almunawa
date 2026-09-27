package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int

	// Menggunakan slice untuk menampung data pemenang karena jumlah pertandingan tidak diketahui
	var daftarPemenang []string

	// Meminta input nama klub
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		// Berhenti jika salah satu atau kedua skor bernilai negatif
		if skorA < 0 || skorB < 0 {
			break
		}

		// Menentukan pemenang dan menyimpannya ke dalam slice
		if skorA > skorB {
			daftarPemenang = append(daftarPemenang, klubA)
		} else if skorB > skorA {
			daftarPemenang = append(daftarPemenang, klubB)
		} else {
			daftarPemenang = append(daftarPemenang, "Draw")
		}

		pertandingan++
	}

	// Menampilkan hasil dari array/slice pemenang
	fmt.Println()
	for i := 0; i < len(daftarPemenang); i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, daftarPemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}
