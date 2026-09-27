package main

import "fmt"

type pemain struct {
	name        string
	gol, assist int
}

const NMAX = 1000

type arrPemain [NMAX]pemain

func main() {
	var dataPemain arrPemain
	var i, j, maxIdx, banyakPemain int
	var namaDepan, namaBelakang string
	var temp pemain

	fmt.Println("Masukkan Data Input : ")

	fmt.Scan(&banyakPemain)
	for i = 0; i < banyakPemain && i < NMAX; i++ {
		fmt.Scan(&namaDepan, &namaBelakang, &dataPemain[i].gol, &dataPemain[i].assist)
		dataPemain[i].name = namaDepan + " " + namaBelakang
	}

	for i = 0; i < banyakPemain-1; i++ {
		maxIdx = i
		for j = i + 1; j < banyakPemain; j++ {
			if dataPemain[j].gol > dataPemain[maxIdx].gol || (dataPemain[j].gol == dataPemain[maxIdx].gol && dataPemain[j].assist > dataPemain[maxIdx].assist) {
				maxIdx = j
			}
		}
		temp = dataPemain[maxIdx]
		dataPemain[maxIdx] = dataPemain[i]
		dataPemain[i] = temp
	}

	fmt.Println("\nHasil Sorting : ")
	for i = 0; i < banyakPemain; i++ {
		fmt.Println(dataPemain[i].name, dataPemain[i].gol, dataPemain[i].assist)
	}
}
