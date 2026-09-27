package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	var i int

	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")

	for i = 0; i < nProv; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	var i, idx int

	idx = 0

	for i = 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}

	return idx
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	var i int
	var idx int = -1

	i = 0
	for i < nProv && idx == -1 {
		if prov[i] == nama {
			idx = i
		}
		i++
	}

	return idx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	var i int
	var hasil float64

	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")

	for i = 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil = (tumbuh[i] + 1) * float64(pop[i])
			fmt.Println(prov[i], int(hasil))
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cari string
	var idxTercepat, idxCari int

	InputData(&prov, &pop, &tumbuh)

	fmt.Scan(&cari)

	idxTercepat = ProvinsiTercepat(tumbuh)

	fmt.Println("\nProvinsi dengan angka pertumbuhan tercepat :", prov[idxTercepat])

	idxCari = IndeksProvinsi(prov, cari)

	fmt.Println("\nData provinsi yang dicari :", prov[idxCari])

	Prediksi(prov, pop, tumbuh)
}
