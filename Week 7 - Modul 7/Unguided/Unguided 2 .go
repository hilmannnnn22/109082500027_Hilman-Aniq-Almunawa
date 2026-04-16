package main

import "fmt"

type angka int
type kata string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

func main() {
	var buku Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("Masukkan judul buku: ")
	fmt.Scan(&buku.judul)

	fmt.Print("Masukkan nama penulis: ")
	fmt.Scan(&buku.penulis)

	fmt.Print("Masukkan penerbit: ")
	fmt.Scan(&buku.penerbit)

	fmt.Print("Masukkan tahun terbit: ")
	fmt.Scan(&buku.tahunTerbit)

	fmt.Print("Masukkan jumlah halaman: ")
	fmt.Scan(&buku.jumlahHalaman)

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", buku.judul)
	fmt.Println("Penulis :", buku.penulis)
	fmt.Println("Penerbit :", buku.penerbit)
	fmt.Println("Tahun Terbit :", buku.tahunTerbit)
	fmt.Println("Jumlah Halaman :", buku.jumlahHalaman)
}
