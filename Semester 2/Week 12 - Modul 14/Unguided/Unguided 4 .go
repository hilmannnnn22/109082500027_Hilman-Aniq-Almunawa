package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("%s, %s, %s, %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := pustaka[i]
		j := i
		// Sorting menurun (descending) berdasarkan rating
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Judul Buku dengan Rating Tertinggi:")
	batas := 5
	if n < 5 {
		batas = n
	}
	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1
	ketemu := false

	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			fmt.Println("Data Buku Ditemukan:")
			fmt.Printf("%s, %s, %s, %d, %d, %d\n", pustaka[tengah].judul, pustaka[tengah].penulis, pustaka[tengah].penerbit, pustaka[tengah].tahun, pustaka[tengah].eksemplar, pustaka[tengah].rating)
			ketemu = true
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1 // Geser kanan karena descending
		} else {
			kiri = tengah + 1 // Geser kiri karena descending
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)

	fmt.Println()
	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	fmt.Println()
	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&ratingCari)
	fmt.Println()
	CariBuku(pustaka, n, ratingCari)
}
