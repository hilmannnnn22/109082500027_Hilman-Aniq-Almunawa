package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, N *int) {
	var i int

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(N)

	for i = 0; i < *N; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}
}

func cariNilaiPertama(T arrayMahasiswa, N int, nim string) int {
	var i int
	var idx int = -1

	i = 0
	for i < N && idx == -1 {
		if T[i].NIM == nim {
			idx = i
		}
		i++
	}

	return idx
}

func cariNilaiTerbesar(T arrayMahasiswa, N int, nim string) int {
	var i, idxMax int

	idxMax = cariNilaiPertama(T, N, nim)

	if idxMax != -1 {
		for i = idxMax + 1; i < N; i++ {
			if T[i].NIM == nim && T[i].nilai > T[idxMax].nilai {
				idxMax = i
			}
		}
	}

	return idxMax
}

func main() {
	var data arrayMahasiswa
	var n int
	var nimCari string
	var idx1, idx2 int

	inputData(&data, &n)

	fmt.Print("\nMasukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	idx1 = cariNilaiPertama(data, n, nimCari)
	idx2 = cariNilaiTerbesar(data, n, nimCari)

	if idx1 != -1 {
		fmt.Println("Nilai pertama dari NIM", nimCari, "adalah", data[idx1].nilai)
		fmt.Println("Nilai terbesar dari NIM", nimCari, "adalah", data[idx2].nilai)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}