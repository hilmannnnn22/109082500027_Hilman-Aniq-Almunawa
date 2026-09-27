package main

import (
	"fmt"
	"math"
)

const NMAX int = 100

type tabel [NMAX]int

// a. Menampilkan keseluruhan isi array
func tampilSemua(t tabel, n int) {
	fmt.Print("Semua elemen array: ")
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
	fmt.Println()
}

// b. Menampilkan elemen dengan indeks ganjil
func tampilGanjil(t tabel, n int) {
	fmt.Print("Elemen indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(t[i], " ")
	}
	fmt.Println()
}

// c. Menampilkan elemen dengan indeks genap
func tampilGenap(t tabel, n int) {
	fmt.Print("Elemen indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(t[i], " ")
	}
	fmt.Println()
}

// d. Menampilkan elemen dengan indeks kelipatan x (fix: skip indeks 0)
func tampilKelipatanX(t tabel, n int, x int) {
	fmt.Printf("Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i != 0 && i%x == 0 {
			fmt.Print(t[i], " ")
		}
	}
	fmt.Println()
}

// e. Menghapus elemen pada indeks tertentu
func hapusElemen(t *tabel, n *int, idx int) {
	for i := idx; i < (*n)-1; i++ {
		t[i] = t[i+1]
	}
	*n--
}

// f. Menampilkan rata-rata
func rataRata(t tabel, n int) float64 {
	sum := 0
	for i := 0; i < n; i++ {
		sum += t[i]
	}
	return float64(sum) / float64(n)
}

// g. Menampilkan standar deviasi
func standarDeviasi(t tabel, n int) float64 {
	mean := rataRata(t, n)
	sumKuadrat := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(t[i]) - mean
		sumKuadrat += selisih * selisih
	}
	return math.Sqrt(sumKuadrat / float64(n))
}

// h. Menampilkan frekuensi suatu bilangan
func frekuensi(t tabel, n int, bil int) int {
	count := 0
	for i := 0; i < n; i++ {
		if t[i] == bil {
			count++
		}
	}
	return count
}

func main() {
	var arr tabel
	var n int

	fmt.Print("Masukkan jumlah elemen array (maks 100): ")
	fmt.Scan(&n)
	if n > NMAX {
		fmt.Println("Jumlah elemen melebihi kapasitas maksimum!")
		return
	}

	fmt.Println("Masukkan", n, "bilangan bulat:")
	for i := 0; i < n; i++ {
		fmt.Printf("  arr[%d] = ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\n========== HASIL ==========")

	// a
	tampilSemua(arr, n)

	// b
	tampilGanjil(arr, n)

	// c
	tampilGenap(arr, n)

	// d
	var x int
	fmt.Print("\nMasukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	tampilKelipatanX(arr, n, x)

	// e
	var idxHapus int
	fmt.Printf("\nMasukkan indeks yang ingin dihapus (0 - %d): ", n-1)
	fmt.Scan(&idxHapus)
	hapusElemen(&arr, &n, idxHapus)
	fmt.Print("Array setelah penghapusan: ")
	tampilSemua(arr, n)

	// f
	fmt.Printf("\nRata-rata: %.2f\n", rataRata(arr, n))

	// g
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(arr, n))

	// h
	var bilCari int
	fmt.Print("\nMasukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bilCari)
	fmt.Printf("Frekuensi %d dalam array: %d kali\n", bilCari, frekuensi(arr, n, bilCari))
}